package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/microports/app/service/api/representation"
	"github.com/microports/app/service/port/pb/port"
	"github.com/microports/app/util/gateway"

	"github.com/emicklei/go-restful"
)

// 20mb
const maxClassesJSONFileSize = 20 * 1024 * 1024

func (rcv *PortGateway) UploadPorts(req *restful.Request, resp *restful.Response) {

	// Read request body and header
	file, fileHeader, err := req.Request.FormFile(uploadFileParameterName)
	if err != nil {
		gateway.HandleError(resp, err, http.StatusBadRequest, representation.ErrInvalidFile)
		return
	}

	// Check file size
	if fileHeader.Size > maxClassesJSONFileSize {
		gateway.HandleError(resp, fmt.Errorf(representation.ErrFileTooLarge), http.StatusBadRequest, representation.ErrFileTooLarge)
		return
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("[gateway-port] upload closing file error: %s", err)
		}
	}()

	// Prepare flow JSON unmarshaling
	var (
		// Chanel for parsed entities
		dataChan = make(chan *port.PortEntity)

		// Chanel for errors
		errChan = make(chan error)

		// Chanel for notification of the end
		doneChan = make(chan struct{})
	)

	go rcv.unmarshalJSONFile(file, dataChan, errChan, doneChan)

	go rcv.consumeParsedPorts(dataChan, errChan, doneChan)

	select {
	case <-doneChan:
		resp.WriteHeader(http.StatusNoContent)

	case err := <-errChan:
		gateway.HandleError(resp, err, http.StatusBadRequest, err.Error())
	}
}

func (rcv *PortGateway) consumeParsedPorts(
	dataChan <-chan *port.PortEntity,
	errChan chan error,
	doneChan <-chan struct{},
) {
	for {
		select {
		case parsedPort := <-dataChan:
			err := rcv.savePort(parsedPort)
			if err != nil {
				log.Println("consumeParsedPorts send error")
				errChan <- err
				return
			}

		case <-errChan:
			log.Println("consumeParsedPorts receive error")
			return

		case <-doneChan:
			log.Println("consumeParsedPorts receive done")
			return
		}
	}
}

func (rcv *PortGateway) unmarshalJSONFile(
	file io.Reader,
	dataChan chan *port.PortEntity,
	errChan chan error,
	doneChan chan struct{},
) {
	defer func() {
		doneChan <- struct{}{}
	}()

	dec := json.NewDecoder(file)

	for {
		select {
		case <-errChan:
			return
		default:
			//	NOP
		}

		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			errChan <- err
			return
		}

		switch token.(type) {
		case json.Delim:
			// NOP
		default:
			if dec.More() {

				var ID string
				var isString bool

				if ID, isString = token.(string); !isString {
					errChan <- fmt.Errorf("JSON hasn't map structure")
					return
				}

				parsedPort := &representation.UploadPortRequest{}

				if err := dec.Decode(&parsedPort); err != nil {
					errChan <- fmt.Errorf("JSON has invalid structure %+v", err)
					return
				}

				dataChan <- transformUploadPortToPortEntity(ID, parsedPort)
			}
		}
	}
}

func (rcv *PortGateway) savePort(entity *port.PortEntity) error {
	_, err := rcv.portService.SavePort(context.Background(), &port.SavePortRequest{
		Port: entity,
	})
	if err != nil {
		return err
	}

	return nil
}

func transformUploadPortToPortEntity(ID string, uploadPort *representation.UploadPortRequest) *port.PortEntity {
	return &port.PortEntity{
		ID:          ID,
		Name:        uploadPort.Name,
		Coordinates: uploadPort.Coordinates,
		City:        uploadPort.City,
		Province:    uploadPort.Province,
		Country:     uploadPort.Country,
		Alias:       uploadPort.Alias,
		Regions:     uploadPort.Regions,
		Timezone:    uploadPort.Timezone,
		Unlocs:      uploadPort.Unlocs,
		Code:        uploadPort.Code,
	}
}
