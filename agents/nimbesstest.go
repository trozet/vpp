

package main

import (
	"fmt"

	"github.com/trozet/vpp/plugins/podmanager/cni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

const serverPort = 9111


func main() {
	log.Info("Connecting to Nimbess")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", serverPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := cni.NewRemoteCNIClient(conn)
	testReq := &cni.CNIRequest{
		ContainerId: "421b7e85432d99593a6989073437d20d2f8947ffc292544e3065e59434fd400c",

	}

	res, err := client.Add(context.Background(), testReq)
	if err != nil || res.Error !="" || res.Result !=0 {
		log.Fatalf("Port Add failed: %v, %v", res, err)
	}
	log.Infof("Port added!: %s", res.Interfaces)

}