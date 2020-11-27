package internal

import (
	"fmt"
	"net"

	"github.com/microports/app/service"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/service/port/persisntence"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type portService struct {
	portRepository persisntence.PortRepository
}

func NewPortService(
	portRepository persisntence.PortRepository,
) service.Service {
	return &portService{
		portRepository: portRepository,
	}
}

func (rcv *portService) Attach(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to acquire address [%s]: %s", address, err)
	}

	s := grpc.NewServer()
	port.RegisterPortServiceServer(s, rcv)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
