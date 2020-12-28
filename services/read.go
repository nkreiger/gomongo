package services

import (
	"app/mongo"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var ReadCollectionPayload = func(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")

	var payload mongo.QueryInputs

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		writeStatus(&w, http.StatusBadRequest, "invalid query payload")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx)
	if err != nil {
		writeStatus(&w, http.StatusBadRequest, err.Error())
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("error disconnecting from client: %v", err)
		}

		cancel()
	}()

	executeParams := mongo.ExecuteParams{
		Context:    ctx,
		Database:   client.Database("testdb"),
		Collection: "testcollection",
	}

	resp, err := mongo.GetEvents(executeParams, payload)

	writeStatus(&w, http.StatusOK, resp)
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