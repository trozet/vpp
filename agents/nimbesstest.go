

package main

import (
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
	docker "github.com/docker/docker/client"

	log "github.com/sirupsen/logrus"
	"github.com/trozet/vpp/plugins/podmanager/cni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const containerId = "vport_test"
const nimbessPort = 9111

func main() {
	log.Info("Connecting to Nimbess")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", nimbessPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	cli, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	container, err := cli.ContainerInspect(context.Background(), containerId)
	if err != nil {
		log.Fatal("Please start container first!")
	}

	log.Infof("containers: %v, err: %v", container.NetworkSettings.SandboxKey, err)

	netConf := &types.NetConf{
		Name: "testNetwork",
	}
	jres, _ := json.Marshal(netConf)
	client := cni.NewRemoteCNIClient(conn)
	testReq := &cni.CNIRequest{
		ContainerId: containerId,
		InterfaceName: "eth0",
		ExtraNwConfig: string(jres),
		NetworkNamespace: container.NetworkSettings.SandboxKey,
	}

	res, err := client.Add(context.Background(), testReq)
	if err != nil || res.Error !="" || res.Result !=0 {
		log.Fatalf("Port Add failed: %v, %v", res, err)
	}
	log.Infof("Port added!: %s", res.Interfaces)

	/*
	dres, derr := client.Delete(context.Background(), testReq)
	if derr != nil || dres.Error !="" || dres.Result !=0 {
		log.Fatalf("Port Delete failed: %v, %v", dres, derr)
	}
	log.Infof("Port deleted!: %s", res.Interfaces)
	*/

}