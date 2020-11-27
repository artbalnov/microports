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

func TestUploadPortsJSONSuccess(t *testing.T) {
	var (
		errResult = &gateway.ErrorResponse{}
		sucResult = &representation.GetPortsResponse{}
	)

	// Get address
	portGatewayURL, err := env.GetVar("GATEWAY_PORT_URL")
	if err != nil {
		t.Fatalf("failed to get target URL: %s", err)
	}

	// Do upload JSON request
	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", "../fixture/correct_data.json").
		SetError(&errResult).
		Post(fmt.Sprintf("%s/api/v1/ports/upload", portGatewayURL))

	if err != nil {
		t.Fatal(err)
	}

	// Assert results
	as := assert.New(t)

	as.Empty(errResult.Message)

	as.NotNil(resp)

	as.Equal(http.StatusNoContent, resp.StatusCode())

	// Do get all ports request
	resp, err = resty.R().
		SetResult(&sucResult).
		SetError(&errResult).
		Get(fmt.Sprintf("%s/api/v1/ports/", portGatewayURL))

	as.Empty(errResult.Message)
	as.NotNil(sucResult)

	as.NotNil(resp)

	as.Equal(http.StatusOK, resp.StatusCode())
}
