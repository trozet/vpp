

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-ini/ini"
	"github.com/golang/protobuf/ptypes"
	"github.com/trozet/vpp/agents/bess_pb"
	"github.com/trozet/vpp/plugins/podmanager/cni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

const Switch = "tcam"
const L2DriverMode = "layer2"

// INI config types
type Network struct {
	Driver		string	`ini:"driver"`
	MacLearn	bool	`ini:"mac_learning"`
	TunnelMode	bool	`ini:"tunnel_mode"`
	FIBSize		int64	`ini:"fib_size"`
}

type NimbessConfig struct {
	BessPort    int			`ini:"bess_port"`
	WorkerCores	[]int		`ini:"worker_cores"`
	NICs		[]string	`ini:"pci_devices"`
	Network
}

// CNI GRPC server
type cniServer struct {
	client bess_pb.BESSControlClient
	ports map[string][]string // maps bess port ids to namespaces
	mu sync.Mutex
}

func (s *cniServer) Add(ctx context.Context, req *cni.CNIRequest) (*cni.CNIReply, error) {
	// hardcode IP for now for testing
	ipAddr := "40.0.0.1/24"

	vportNetns := &bess_pb.VPortArg_Netns{
		Netns: req.NetworkNamespace,
	}
	vportArgs := &bess_pb.VPortArg{
		Ifname: req.InterfaceName,
		Cpid: vportNetns,
		IpAddrs: []string{ipAddr},
	}
	any, err := ptypes.MarshalAny(vportArgs)
	if err != nil {
		return &cni.CNIReply{}, err
	}

	portRequest := &bess_pb.CreatePortRequest{
		Name: req.InterfaceName,
		Driver: "vport",
		Arg: any,
	}
	res, err := s.client.CreatePort(ctx, portRequest)
	if err != nil {
		return &cni.CNIReply{}, err
	}

	s.mu.Lock()
	s.ports[req.NetworkNamespace] = append(s.ports[req.NetworkNamespace], req.InterfaceName)
	s.mu.Unlock()
	var repRes uint32 = 0
	if res.Error != nil {
		repRes = 1
	}

	cniIfIp := &cni.CNIReply_Interface_IP{
		Version: cni.CNIReply_Interface_IP_IPV4,
		Address: ipAddr,
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
	var portIdx int
	portReq := &bess_pb.DestroyPortRequest{
		Name: req.InterfaceName,
	}
	for idx, p := range s.ports[req.NetworkNamespace] {
		if p == req.InterfaceName {
			portIdx = idx
			break
		}
	}

	if portIdx == nil {
		return &cni.CNIReply{Error: "Port not found"}, errors.New("port not found")
	}
	res, err := s.client.DestroyPort(ctx, portReq)

	if err != nil {
		return &cni.CNIReply{}, err
	}

	s.mu.Lock()
	s.ports[req.NetworkNamespace] = append(s.ports[req.NetworkNamespace][:idx],
		s.ports[req.NetworkNamespace][idx+1:]...)
	s.mu.Unlock()

	if res.Error != nil {
		return &cni.CNIReply{
			Result: 1,
			Error: res.Error.Errmsg,
		}, nil
	} else {
		return &cni.CNIReply{
			Result: 0,
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
		log.Infof("Switch created: %s\n, %s", cRes.Name, cRes.Error)
	}

}


func main() {
	// TODO setup cmd line args for cfg file path, etc
	// For now assume cfg file in same path as binary
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(dir, "nimbess.ini")

	nimbessConfig := initConfig(configFile)
	log.Info("Connecting to BESS")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", nimbessConfig.BessPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := bess_pb.NewBESSControlClient(conn)
	verResponse, err := client.GetVersion(context.Background(), &bess_pb.EmptyRequest{})
	if err != nil || verResponse.Error !=nil {
		log.Fatalf("Could not get version: %v", verResponse.Error)
	}
	log.Infof("BESS connected with version: %s\n", verResponse.Version)

	if nimbessConfig.Driver == L2DriverMode {
		log.Info("Building L2 switch")
		initSwitch(client, nimbessConfig.FIBSize)
	} else {
		log.Fatalf("Unsupported driver specified in configuration file: %s", nimbessConfig.Driver)
	}


	// Implement gRPC listener for CNI calls (can use contiv_cni.go and podmanager protos)
	// Port add handler
	

}