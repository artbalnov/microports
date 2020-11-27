package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/microports/app/util/env"
	"github.com/microports/app/util/gateway"
	"github.com/stretchr/testify/assert"

	"gopkg.in/resty.v1"
)

func TestUploadPortsJSONFail(t *testing.T) {
	var (
		errResult = &gateway.ErrorResponse{}
	)

	// Get address
	portGatewayURL, err := env.GetVar("GATEWAY_PORT_URL")
	if err != nil {
		t.Fatalf("failed to get target URL: %s", err)
	}

	// Do upload JSON request
	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", "../fixture/wrong_data.json").
		SetError(&errResult).
		Post(fmt.Sprintf("%s/api/v1/ports/upload", portGatewayURL))

	if err != nil {
		t.Fatal(err)
	}
	// Assert results
	as := assert.New(t)

	as.NotEmpty(errResult.Message)

	as.Equal(http.StatusBadRequest, resp.StatusCode())
}
