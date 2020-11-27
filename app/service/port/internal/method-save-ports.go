package internal

import (
	"context"

	"github.com/microports/app/service/port/errors"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/service/port/persisntence/model"
)

func (rcv *portService) SavePorts(ctx context.Context, req *port.SavePortsRequest) (*port.SavePortsResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, errors.Internal(err)
	}

	// Validate length
	if len(req.GetItems()) == 0 {
		return &port.SavePortsResponse{}, nil
	}

	// Map request entities to business models
	ports := make([]*model.Port, len(req.GetItems()))

	for i, p := range req.GetItems() {
		ports[i] = &model.Port{
			ID:          p.GetID(),
			Name:        p.GetName(),
			Code:        p.GetCode(),
			Alias:       p.GetAlias(),
			Unlocs:      p.GetUnlocs(),
			Country:     p.GetCountry(),
			Regions:     p.GetRegions(),
			Province:    p.GetProvince(),
			City:        p.GetCity(),
			Coordinates: p.GetCoordinates(),
			Timezone:    p.GetTimezone(),
		}
	}

	// Save models
	err := rcv.portRepository.SaveAll(ports)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return &port.SavePortsResponse{}, nil
}
