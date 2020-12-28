package services

import (
	"encoding/json"
	"log"
	"net/http"
)

type status struct {
	Message string `json:"message"`
}

// Health returns the health of the API
var Health = func(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Content-Type", "application/json")

	log.Println("Health probe received")

	output := status{
		Message: "Health Probe Successful",
	}

	data, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		log.Printf("error encoding json: %s", data)
	}

	_, err = w.Write(data)
	if err != nil {
		log.Printf("error writing response to user: %v", err)
	}
}