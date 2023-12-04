package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alessandra1408/crud-golang/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Some error occurred while get register: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
