package main

import (
	"context"
	"log"

	"github.com/codingpot/server-client-template-go/pkg/pbs"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pbs.NewDummyServiceClient(conn)

	initResponse, err := c.AgentInit(context.Background(), &pbs.InitRequest{Body: "I am ready"})

	if err != nil {
		log.Fatalf("Error when calling AgentInit: %s", err)
	}
	log.Printf("Response from server: %s", initResponse)
	configResponse, err := c.GetAgentConfig(context.Background(), &pbs.ConfigRequest{ComputeId: initResponse.ComputeId})

	if err != nil {
		log.Fatalf("Error when calling GetAgentConfig: %s", err)
	}
	log.Printf("Response from server: %s", configResponse)
	for i := 0; i < int(configResponse.Capacity); i++ {
		log.Printf("Lunch agent: %d", 8082+i)
	}
}
