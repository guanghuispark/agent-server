package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/codingpot/server-client-template-go/pkg/pbs"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.86.96.236:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pbs.NewDummyServiceClient(conn)
	initResponse, err := c.AgentInit(context.Background(), &pbs.InitRequest{PublicIP: getInstanceIP()})

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
		log.Printf("Launch agent: %d", 8082+i)
		executeAgent(strconv.Itoa(8082 + i))
	}
}

func executeAgent(port string) {
	cmd := exec.Command("./ugp_agent_d.exe", port)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func getInstanceIP() string {
	resp, err := http.Get("http://metadata.tencentyun.com/meta-data/public-ipv4")
	if err != nil {
		fmt.Println(err)
		return "Get public IP failed"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		return string(body)
	}

	return "Get public IP failed"
}

func getInstanceID() string {
	resp, err := http.Get("http://metadata.tencentyun.com/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return "Get instance id failed"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		return string(body)
	}

	return "Get instance id failed"
}
