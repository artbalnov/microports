package internal

import (
	"net/http"

	"github.com/microports/app/service"
	"github.com/microports/app/service/port/pb/port"

	"github.com/emicklei/go-restful"
	"github.com/facebookgo/grace/gracehttp"
)

type PortGateway struct {
	webContainer *restful.Container

	portService port.PortServiceClient
}

func (rcv *PortGateway) Attach(address string) error {
	// Starting restful api http listener
	return gracehttp.Serve(
		&http.Server{
			Addr:    address,
			Handler: rcv.webContainer,
		},
	)
}

func NewPortGatewayService(portServiceClient port.PortServiceClient) service.Service {
	gatewayService := &PortGateway{
		webContainer: restful.NewContainer(),
		portService:  portServiceClient,
	}

	gatewayService.RegisterHandler()

	return gatewayService
}
