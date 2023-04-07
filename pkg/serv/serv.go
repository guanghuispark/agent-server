package serv

import (
	"context"
	"fmt"

	"github.com/codingpot/server-client-template-go/pkg/pbs"
)

type Server struct {
	pbs.UnimplementedDummyServiceServer
}

func (s Server) AgentInit(ctx context.Context, in *pbs.InitRequest) (*pbs.InitResponse, error) {
	fmt.Printf("[gRPC server] AgentInit func is called \n")
	return &pbs.InitResponse{
		ComputeId: "1234567",
	}, nil
}

func (s Server) GetAgentConfig(ctx context.Context, in *pbs.ConfigRequest) (*pbs.ConfigResponse, error) {
	fmt.Printf("[gRPC server] GetAgentConfig func is called\n ")
	return &pbs.ConfigResponse{
		Capacity: 6,
	}, nil
}
