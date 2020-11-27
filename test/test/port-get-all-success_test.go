package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/microports/app/service/api/representation"
	"github.com/microports/app/util/env"
	"github.com/microports/app/util/gateway"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

func TestGetAllPortSuccess(t *testing.T) {
	var (
		errResult = &gateway.ErrorResponse{}
		sucResult = &representation.GetPortsResponse{}
	)

	// Get address
	portGatewayURL, err := env.GetVar("GATEWAY_PORT_URL")
	if err != nil {
		t.Fatalf("failed to get target URL: %s", err)
	}

	// Do get all ports request
	resp, err := resty.R().
		SetResult(&sucResult).
		SetError(&errResult).
		Get(fmt.Sprintf("%s/api/v1/ports/", portGatewayURL))

	// Assert results
	as := assert.New(t)

	as.Empty(errResult.Message)
	as.NotNil(sucResult)

	as.NotNil(resp)

	as.Equal(http.StatusOK, resp.StatusCode())
}
