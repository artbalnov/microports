package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/microports/app/service/api/representation"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/util/gateway"

	"github.com/emicklei/go-restful"
)

func (rcv *PortGateway) GetAllPorts(req *restful.Request, resp *restful.Response) {
	// Get saved ports
	savedPorts, err := rcv.portService.GetAllPorts(context.Background(), &port.GetAllPortsRequest{})
	if err != nil {
		gateway.HandleError(resp, err, http.StatusInternalServerError, representation.ErrInternal)
		return
	}

	// Compose response
	allPortsResponse := &representation.GetPortsResponse{
		Ports: make([]*representation.GetPortResponse, len(savedPorts.GetItems())),
	}

	for i, p := range savedPorts.GetItems() {
		allPortsResponse.Ports[i] = &representation.GetPortResponse{
			ID:          p.GetID(),
			Name:        p.GetName(),
			Coordinates: p.GetCoordinates(),
			City:        p.GetCity(),
			Province:    p.GetProvince(),
			Country:     p.GetCountry(),
			Alias:       p.GetAlias(),
			Regions:     p.GetRegions(),
			Timezone:    p.GetTimezone(),
			Unlocs:      p.GetUnlocs(),
			Code:        p.GetCode(),
		}
	}

	if err := resp.WriteHeaderAndEntity(http.StatusOK, allPortsResponse); err != nil {
		log.Printf("[handler-get-all-ports] occured error: %+v", err)
	}
}
