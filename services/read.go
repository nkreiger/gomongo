package services

import (
	"app/mongo"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func writeStatus(writer *http.ResponseWriter, statusCode int, response interface{}) {
	(*writer).WriteHeader(statusCode)

	data, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		log.Printf("serialization error: %v", data)
	}

	_, err = (*writer).Write(data)
	if err != nil {
		log.Printf("error writing response %v", err)
	}
}