package errors

import (
	"log"

	"github.com/microports/app/service/port/pb/port"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Internal(err error) error {
	log.Printf("[service-port] internal error: %s", err)
	return status.Error(codes.Code(port.PortServiceErrorCode_Internal), err.Error())
}

func InvalidRequest(err error) error {
	log.Printf("[service-port] invalid request: %s", err)
	return status.Error(codes.Code(port.PortServiceErrorCode_InvalidRequest), err.Error())
}
