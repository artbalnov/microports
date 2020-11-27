package gateway

import (
	"encoding/json"
	"log"

	"github.com/emicklei/go-restful"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func HandleError(response *restful.Response, err error, statusCode int, message string) {
	log.Printf("Handle error: %+v", err)

	errEntity := &ErrorResponse{
		Message: message,
	}

	jsonString, err := json.Marshal(errEntity)

	if err != nil {
		log.Println(err)
	}

	response.AddHeader("Content-Type", restful.MIME_JSON)
	err = response.WriteErrorString(statusCode, string(jsonString))

	if err != nil {
		log.Println(err)
	}
}
