

package main

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/trozet/vpp/agents/bess_pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

const Switch = "tcam"

func initSwitch(client bess_pb.BESSControlClient, size int64) {
	modRequest := &bess_pb.GetModuleInfoRequest{
		Name: Switch,
	}
	res, err := client.GetModuleInfo(context.Background(),modRequest)
	if err == nil && res.GetName() == Switch {
		log.Infof("Switch: %s already exists", Switch)
		return
	}
	// todo need to enable learn mode
	l2Args := &bess_pb.L2ForwardArg{
		Size: size,
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
	log.Info("Connecting to BESS")
	conn, err := grpc.Dial("localhost:10514", grpc.WithInsecure())
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

	log.Info("Building L2 switch")
	initSwitch(client, 1024)


	// Implement gRPC listener for CNI calls (can use contiv_cni.go and podmanager protos)
	// Port add handler

}