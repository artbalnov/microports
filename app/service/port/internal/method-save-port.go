package internal

import (
	"context"

	"github.com/microports/app/service/port/errors"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/service/port/persisntence/model"
)

func (rcv *portService) SavePort(ctx context.Context, req *port.SavePortRequest) (*port.SavePortResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, errors.Internal(err)
	}

	// Map request entity to business model
	if err := rcv.portRepository.Save(&model.Port{
		ID:          req.GetPort().GetID(),
		Name:        req.GetPort().GetName(),
		Code:        req.GetPort().GetCode(),
		Alias:       req.GetPort().GetAlias(),
		Unlocs:      req.GetPort().GetUnlocs(),
		Country:     req.GetPort().GetCountry(),
		Regions:     req.GetPort().GetRegions(),
		Province:    req.GetPort().GetProvince(),
		City:        req.GetPort().GetCity(),
		Coordinates: req.GetPort().GetCoordinates(),
		Timezone:    req.GetPort().GetTimezone(),
	}); err != nil {
		return nil, errors.Internal(err)
	}

	return &port.SavePortResponse{}, nil
}
