package services

import (
	"app/mongo"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var Connection = func(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx)
	if err != nil {
		writeStatus(&w, http.StatusBadRequest, err.Error())
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("error disconnecting client: %v", err)
		}

		cancel()
	}()

	// write back okay for good connection
	writeStatus(&w, http.StatusOK, nil)
}

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