package internal

import (
	"net/http"

	"github.com/microports/app/service/api/representation"
	"github.com/microports/app/util/gateway"

	"github.com/emicklei/go-restful"
)

const (
	uploadFileParameterName = "file"
)

func (rcv *PortGateway) RegisterHandler() {
	ws := &restful.WebService{}

	ws.Path("/api/v1/ports").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/upload").
		To(rcv.UploadPorts).
		Operation("UploadPortsFromFile").
		Consumes("multipart/form-data").
		Param(
			ws.FormParameter(uploadFileParameterName, "Upload file").DataType("file"),
		).
		Returns(http.StatusNoContent, http.StatusText(http.StatusOK), nil).
		Returns(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), gateway.ErrorResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), gateway.ErrorResponse{}))

	ws.Route(ws.GET("/").
		To(rcv.GetAllPorts).
		Operation("GetPorts").
		Writes(representation.GetPortsResponse{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), representation.GetPortsResponse{}).
		Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), gateway.ErrorResponse{}))

	rcv.webContainer.Add(ws)
}
