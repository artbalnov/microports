package api

import (
	"fmt"
	"github.com/microports/app/service"
	"log"

	"github.com/microports/app/service/api/internal"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/util/env"
	"google.golang.org/grpc"
)

const (
	ServiceName = "gateway-port"
)

func Factory() (service.Service, error) {
	// Init service client
	portServiceAddress, err := env.GetVar(env.ServicePortAddress)
	if err != nil {
		log.Fatal("[service-api] can't fetch port service address")
	}

	portService, err := buildPortServiceClient(portServiceAddress)
	if err != nil {
		log.Fatal("[service-api] can't fetch port service address")
	}

	log.Println("[service-api] finish initializing")

	return internal.NewPortGatewayService(portService), nil
}

func buildPortServiceClient(address string) (port.PortServiceClient, error) {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("[service-api] can't connect: %s", err)
	}

	return port.NewPortServiceClient(conn), nil
}
