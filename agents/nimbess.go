

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	any2 "github.com/golang/protobuf/ptypes/any"
	"io"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	cnitypes "github.com/containernetworking/cni/pkg/types"
	"github.com/go-ini/ini"
	"github.com/golang/protobuf/ptypes"
	"github.com/trozet/vpp/agents/bess_pb"
	"github.com/trozet/vpp/plugins/podmanager/cni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

const (
	Switch = "tcam"
	L2DriverMode = "layer2"
	serverPort = 9111
	CNI_INCOMPATIBLE = 1
	INVALID_NETWORK_CONFIG = 2
	CONTAINER_NOT_EXIST = 3
	CNI_TRY_LATER = 11
	BESS_FAILURE = 101
	CNI_OK = 0
	PORT = "Port"
	MODULE = "Module"
	PortInc = "PortInc"
	PortOut = "PortOut"
)

var SUPPORTED_OBJECTS = [...]string{PORT, MODULE}

// Map of port name in BESS to gate number
var TcamPortMap = make(map[string]uint64)
var ipAddr = net.IP{40, 0, 0, 0}


// INI config types
type Network struct {
	Driver		string	`ini:"driver"`
	MacLearn	bool	`ini:"mac_learning"`
	TunnelMode	bool	`ini:"tunnel_mode"`
	FIBSize		int64	`ini:"fib_size"`
}

type NimbessConfig struct {
	BessPort	int			`ini:"bess_port"`
	WorkerCores	[]int		`ini:"worker_cores"`
	NICs		[]string	`ini:"pci_devices"`
	Network
}

type PodID struct {
	Network			string
	ContainerID		string
	Interface		string
}

// LocalPod represents a locally deployed pod (locally = on this node).
type LocalPod struct {
	Network				string
	ContainerID      	string
	NetworkNamespace	string
	MacAddr				string
	IPAddr				string
	Interface			string
	DataPlanePort		string
}

// LocalPods is a map of port-ID -> Pod info.
type LocalPods map[PodID]*LocalPod

// CNI GRPC server
type cniServer struct {
	client bess_pb.BESSControlClient
	pods LocalPods
	mu *sync.Mutex
}

func (s *cniServer) Add(ctx context.Context, req *cni.CNIRequest) (*cni.CNIReply, error) {
	// hardcode IP for now for testing
	ipAddr[3]++

	/* Fix CNI proto as it is not requiring a bunch of spec required fields
	like Network name with Network Configuration. For now just use ExtraNwConfig string
	and unmarshal it
	*/
	netConf := cnitypes.NetConf{}
	if err := json.Unmarshal([]byte(req.ExtraNwConfig), &netConf); err != nil {
		return &cni.CNIReply{Result: INVALID_NETWORK_CONFIG}, fmt.Errorf("failed to load netconf: %v", err)
	}

	vportArg := &bess_pb.VPortArg_Netns{
		Netns: req.NetworkNamespace,
	}
	vportArgs := &bess_pb.VPortArg{
		Ifname: req.InterfaceName,
		Cpid: vportArg,
		IpAddrs: []string{fmt.Sprintf("%s/24", ipAddr.String())},
	}
	any, err := ptypes.MarshalAny(vportArgs)
	if err != nil {
		log.Errorf("Failure to serialize vport args: %v", vportArgs)
		return &cni.CNIReply{Result: CNI_TRY_LATER}, err
	}

	id := PodID {
		ContainerID: req.ContainerId,
		Network: netConf.Name,
		Interface: req.InterfaceName,
	}
	_, ok := s.pods[id]
	if ok {
		log.Error("Pod already exists, invalid CNI ADD call")
		return &cni.CNIReply{Result:CNI_INCOMPATIBLE}, errors.New("pod already exists, invalid CNI ADD call")
	}

	// Hardcode to Kernel veth port for now
	// Need to make this port name unique here, so use short NS name and port name
	netNsSlice := strings.Split(req.NetworkNamespace, "/")
	netNs := netNsSlice[len(netNsSlice)-1]
	bessPort := fmt.Sprintf("port_%s_%s", netNs, req.InterfaceName)
	portRequest := &bess_pb.CreatePortRequest{
		Name: bessPort,
		Driver: "VPort",
		Arg: any,
	}
	// Lock to prevent syncTcam thread from removing port before it is in DB
	s.mu.Lock()
	res, err := s.client.CreatePort(ctx, portRequest)
	if err != nil || res.Error.Errmsg != "" {
		log.Errorf("Failure to create port with Bess: %v, %v", res, err)
		s.mu.Unlock()
		return &cni.CNIReply{Result: BESS_FAILURE, Error: res.Error.Errmsg}, err
	}
	log.Infof("BESS returned NAME: %s", res.Name)
	var gate uint64
	// select gate to use for this port
	// check if gate assignment already exists for this port
	if gate, ok = TcamPortMap[bessPort]; !ok {
		gate = uint64(len(TcamPortMap) + 1)
	}
	TcamPortMap[bessPort] = gate
	s.mu.Unlock()

	// Hardcode wired to L2 TCAM for now...
	err = wirePort(s.client, s.mu, bessPort, Switch, Switch)
	log.Info(err)
	if err != nil || res.Error.Errmsg != "" {
		log.Errorf("Failure to wire port to pipeline with Bess: %v, %v", res, err)
		// handle clean up and delete port
		return &cni.CNIReply{Result: BESS_FAILURE, Error: res.Error.Errmsg}, err
	}
	log.Infof("Wired port %s to pipeline", bessPort)
	s.mu.Lock()
	s.pods[id] = &LocalPod{ContainerID: req.ContainerId, NetworkNamespace: req.NetworkNamespace,
		MacAddr: res.MacAddr, IPAddr: ipAddr.String(), Interface: req.InterfaceName, Network: netConf.Name,
		DataPlanePort: res.Name}
	s.mu.Unlock()
	var repRes uint32 = 0
	if res.Error.Errmsg != "" {
		repRes = 1
	}

	cniIfIp := &cni.CNIReply_Interface_IP{
		Version: cni.CNIReply_Interface_IP_IPV4,
		Address: ipAddr.String(),
	}

	cniIf := &cni.CNIReply_Interface{
		Name: res.Name,
		Mac: res.MacAddr,
		IpAddresses: []*cni.CNIReply_Interface_IP{cniIfIp},
	}

	return &cni.CNIReply{
		Result: repRes,
		Error: res.Error.Errmsg,
		Interfaces: []*cni.CNIReply_Interface{cniIf},
	}, nil
}

func (s *cniServer) Delete(ctx context.Context, req *cni.CNIRequest) (*cni.CNIReply, error) {

	netConf := cnitypes.NetConf{}
	if err := json.Unmarshal([]byte(req.ExtraNwConfig), &netConf); err != nil {
		return &cni.CNIReply{Result: INVALID_NETWORK_CONFIG}, fmt.Errorf("failed to load netconf: %v", err)
	}
	podId := PodID{
		Network: netConf.Name,
		ContainerID: req.ContainerId,
		Interface: req.InterfaceName,
	}

	pod, ok := s.pods[podId]

	if !ok {
		log.Warning("Could not find valid port for delete request: +%v", req)
		return &cni.CNIReply{Result: CNI_OK}, nil
	}
	port := pod.DataPlanePort
	portReq := &bess_pb.DestroyPortRequest{
		Name: port,
	}

	res, err := s.client.DestroyPort(ctx, portReq)
	if err != nil {
		log.Errorf("Failed to delete BESS interface: %s, error: %v", port, err)
	}

	s.mu.Lock()
	delete(s.pods, podId)
	s.mu.Unlock()

	if res.Error != nil {
		return &cni.CNIReply{
			Result: 1,
			Error: res.Error.Errmsg,
		}, nil
	} else {
		return &cni.CNIReply{
			Result: CNI_OK,
		}, nil
	}
}

func initConfig(cfgPath string) *NimbessConfig {
	cfg := &NimbessConfig {
		BessPort: 10514,
		Network: Network{
			Driver: L2DriverMode,
			MacLearn: true,
			TunnelMode: false,
			FIBSize: 1024,
		},
	}

	err := ini.MapTo(cfg, cfgPath)
	if err != nil {
		log.Warningf("Unable to parse Nimbess config file: %v\n", err)
	} else {
		log.Infof("Configuration parsed as: +%v", cfg)
	}
	return cfg
}

// Finds related modules for a port
func getPortModules(client bess_pb.BESSControlClient, port string) []*bess_pb.GetModuleInfoResponse {
	var modules []*bess_pb.GetModuleInfoResponse
	// port modules to check
	ports := []string{fmt.Sprintf("%s_inc", port),fmt.Sprintf("%s_out",port)}
	for _, pName := range ports {
		mRes, err := client.GetModuleInfo(context.Background(), &bess_pb.GetModuleInfoRequest{Name: pName})
		if err == nil  && mRes.GetError().GetCode() == 0 {
			modules = append(modules, mRes)
		}
	}

	return modules
}

// Disconnects modules that have outgoing gates and deletes them. Returns last error hit during delete.
func deleteModules(client bess_pb.BESSControlClient, modules []*bess_pb.GetModuleInfoResponse) error {
	var dError error
	for _, module := range modules {
		res, err := client.DestroyModule(context.Background(), &bess_pb.DestroyModuleRequest{Name: module.Name})
		if err != nil {
			dError = err
			log.Errorf("Failed to delete module: %v", err)
		} else if res.GetError().GetCode() != 0 {
			dError = errors.New(res.GetError().GetErrmsg())
			log.Errorf("Failed to delete module: %v", dError)
		} else {
			log.Infof("Module deleted: %s", module.Name)
		}
	}
	return dError
}

func syncTcam(client bess_pb.BESSControlClient, mu *sync.Mutex) {
	log.Info("Starting TCAM sync thread")
	for {
		res, err := client.ListPorts(context.Background(), &bess_pb.EmptyRequest{})
		if err == nil && res.Error.GetCode() == 0 {
			for _, port := range res.GetPorts() {
				// check if port is in our database
				mu.Lock()
				_, ok := TcamPortMap[port.Name]
				mu.Unlock()
				if !ok {
					// find  and delete modules associated with port
					deleteModules(client, getPortModules(client, port.Name))
					eRes, err := client.DestroyPort(context.Background(),
						&bess_pb.DestroyPortRequest{Name: port.Name})
					if err == nil && eRes.GetError().GetCode() == 0 {
						log.Infof("Port %s cleaned up during sync", port.Name)
					}
				}

			}
		}
	}
}

func initSwitch(client bess_pb.BESSControlClient, size int64) {
	modRequest := &bess_pb.GetModuleInfoRequest{
		Name: Switch,
	}
	res, err := client.GetModuleInfo(context.Background(),modRequest)
	if err == nil && res.GetName() == Switch {
		log.Infof("Switch: %s already exists", Switch)
		return
	}
	l2Args := &bess_pb.L2ForwardArg{
		Size: size,
		Learn: true,
	}
	any, err := ptypes.MarshalAny(l2Args)
	l2module := &bess_pb.CreateModuleRequest{
		Name:   Switch,
		Mclass: "L2Forward",
		Arg:    any,
	}
	cRes, err := client.CreateModule(context.Background(), l2module)
	if err != nil || cRes.Error.GetCode() !=0 {
		log.Fatalf("Failed to create switch: %s", cRes.Error)
	} else {
		log.Infof("Switch created: %s, %s", cRes.Name, cRes.Error)
	}

}

/* Checks if BESS resource exists. oType is BESS object type which must be one of SUPPORTED_OBJECTS types
 */
func objectExists(client bess_pb.BESSControlClient, object string, oType string) (bool, error) {
	objectSupported := false
	for _, suppObject := range SUPPORTED_OBJECTS {
		if oType == suppObject {
			objectSupported = true
			break
		}
	}

	if !objectSupported {
		return false, fmt.Errorf("unsupported object type: %s", oType)
	}

	moduleName := strings.Join([]string{"List", oType, "s"}, "")
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(context.Background())
	inputs[1] = reflect.ValueOf(&bess_pb.EmptyRequest{})
	method := reflect.ValueOf(client).MethodByName(moduleName)
	retVals := method.Call(inputs)

	if retVals[1].Interface() != nil {
		log.Error("Error querying for object: %s, %v", object, retVals[1].Interface().(error))
		return false, retVals[1].Interface().(error)
	}

	resp := retVals[0]
	respType := resp.Type()
	switch (respType.String()) {
	case "*bess_pb.ListModulesResponse":
		modules := resp.Interface().(*bess_pb.ListModulesResponse).GetModules()
		for _, module := range modules {
			if module.Name == object {
				return true, nil
			}
		}
	case "*bess_pb.ListPortsResponse":
		ports := resp.Interface().(*bess_pb.ListPortsResponse).GetPorts()
		for _, port := range ports {
			if port.Name == object {
				return true, nil
			}
		}
	}

	return false, nil
}


//Creates PortInc/PortOut modules and connects a port RX/TX to ingress and egress modules instances
func wirePort(client bess_pb.BESSControlClient, mutex *sync.Mutex, port string, iModule string, eModule string) error {

	if portExists, err := objectExists(client, port, PORT); err != nil {
		return err
	} else if !portExists {
		return fmt.Errorf("port %s not found in BESS", port)
	}

	for _, modObject := range []string{iModule, eModule} {
		if modExists, err := objectExists(client, modObject, MODULE); err != nil {
			return err
		} else if !modExists {
			return fmt.Errorf("module %s not found in BESS", modObject)
		}
	}

	// Everything exists, now create Port Ingress/Egress modules
	for _,portType := range []string{PortInc, PortOut} {
		var any *any2.Any
		var portName string
		var mErr error
		if portType == PortInc {
			portArgs := &bess_pb.PortIncArg{
				Port: port,
			}
			portName = strings.Join([]string{port, "inc"}, "_")
			any, mErr = ptypes.MarshalAny(portArgs)
			if mErr != nil {
				log.Errorf("Failure to serialize port args: %v", portArgs)
				return mErr
			}
		} else {
			portArgs := &bess_pb.PortOutArg{
				Port: port,
			}
			portName = strings.Join([]string{port, "out"}, "_")
			any, mErr = ptypes.MarshalAny(portArgs)
			if mErr != nil {
				log.Errorf("Failure to serialize port args: %v", portArgs)
				return mErr
			}
		}

		portReq := &bess_pb.CreateModuleRequest{
			Name:   portName,
			Mclass: portType,
			Arg:    any,
		}

		cRes, err := client.CreateModule(context.Background(), portReq)
		if err != nil || cRes.Error.GetCode() != 0 {
			log.Errorf("Failed to create port: %s", cRes.GetError().Errmsg)
			return err
		} else {
			log.Infof("Port created: %s\n, %s", cRes.Name, cRes.Error)
		}

		mutex.Lock()
		gate := TcamPortMap[port]
		mutex.Unlock()
		// Now wire to module
		var connReq *bess_pb.ConnectModulesRequest
		if portType == PortInc {
			connReq = &bess_pb.ConnectModulesRequest{
				M1: portName,
				M2: iModule,
				Igate: gate,
			}
		} else {
			connReq = &bess_pb.ConnectModulesRequest{
				M1: eModule,
				M2: portName,
				Ogate: gate,
			}
		}

		conRes, err := client.ConnectModules(context.Background(), connReq)
		if err != nil {
			log.Errorf("Failed to connect modules: %s", err.Error())
			return err
		} else if conRes.Error.GetCode() != 0  {
			log.Errorf("Failed to connect modules: %s", conRes.GetError().Errmsg)
			return errors.New(conRes.GetError().Errmsg)
		} else {
			log.Infof("Connection created via request: %v", connReq)
		}
	}
	return nil
}

func main() {
	mu := &sync.Mutex{}
	// TODO setup cmd line args for cfg file path, etc
	// For now assume cfg file in same path as binary
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(dir, "nimbess.ini")

	nimbessConfig := initConfig(configFile)

	logFile, e := os.OpenFile("/tmp/nimbess.log", os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if e != nil {
		panic(e)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.Info("Connecting to BESS")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", nimbessConfig.BessPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := bess_pb.NewBESSControlClient(conn)
	verResponse, err := client.GetVersion(context.Background(), &bess_pb.EmptyRequest{})
	if err != nil || verResponse.Error !=nil {
		log.Fatalf("Could not get version: %v, %v", verResponse.GetError(), err)
	}
	log.Infof("BESS connected with version: %s", verResponse.Version)

	if nimbessConfig.Driver == L2DriverMode {
		log.Info("Building L2 switch")
		initSwitch(client, nimbessConfig.FIBSize)
		go syncTcam(client, mu)
	} else {
		log.Fatalf("Unsupported driver specified in configuration file: %s", nimbessConfig.Driver)
	}
	// TODO(trozet) setup core workers
	// gRPC listener for CNI calls (can use contiv_cni.go and podmanager protos)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Info("Starting gRPC server")
	grpcServer := grpc.NewServer()
	cni.RegisterRemoteCNIServer(grpcServer, &cniServer{client:client, pods: make(LocalPods), mu: mu})
	grpcServer.Serve(lis)

}