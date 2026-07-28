package handlers

import (
	"encoding/json"
	"go-shortlinks/database"
	"go-shortlinks/models"
	"log"
	"net/http"
)

func GetAllURLs(w http.ResponseWriter, r *http.Request) {

	if !IsValidMethod(w, r, http.MethodGet) {
		return
	}
	var urls []models.URL

	result := database.DB.Find(&urls)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		log.Printf("Error getting URLs: %v", result.Error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
