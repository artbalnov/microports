package internal

import (
	"context"

	"github.com/microports/app/service/port/errors"
	"github.com/microports/app/service/port/pb/port"
)

func (rcv *portService) GetAllPorts(ctx context.Context, req *port.GetAllPortsRequest) (*port.GetAllPortsResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, errors.Internal(err)
	}

	// Get all saved ports
	savedPorts, err := rcv.portRepository.GetAll()
	if err != nil {
		return nil, errors.Internal(err)
	}

	// Map business models to response entities
	responsePorts := &port.GetAllPortsResponse{
		Items: make([]*port.PortEntity, len(savedPorts)),
	}

	for i, p := range savedPorts {
		responsePorts.Items[i] = &port.PortEntity{
			ID:          p.ID,
			Name:        p.Name,
			Code:        p.Code,
			Alias:       p.Alias,
			Unlocs:      p.Unlocs,
			Country:     p.Country,
			Regions:     p.Regions,
			Province:    p.Province,
			City:        p.City,
			Coordinates: p.Coordinates,
			Timezone:    p.Timezone,
		}
	}

	return responsePorts, nil
}
