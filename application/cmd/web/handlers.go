package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// GET handler: Simply returns a JSON message
func (app *application) getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := MessageResponse{Message: "This is a GET request", Status: http.StatusOK}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (app *application) postHandler(w http.ResponseWriter, r *http.Request) {
	var input map[string]any
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]any{
		"received": input,
		"status":   http.StatusCreated,
	})
	if err != nil {
		return
	}
	jsonData, _ := json.Marshal(input)
	log.Printf("Received POST request: %s", string(jsonData))
}

func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	response := MessageResponse{Message: "This is a DELETE request", Status: http.StatusNoContent}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
