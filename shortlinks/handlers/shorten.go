package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"go-shortlinks/database"
	"go-shortlinks/models"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	action := "insert"
	if r.Method == http.MethodPut {
		action = "update"
	} else if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url models.URL

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding URL: %v", err)
		return
	}
	url.URL = base64.StdEncoding.EncodeToString([]byte(url.URL))

	var existingURL models.URL
	if action == "update" { // action is PUT
		handle := database.DB.Where("handle = ?", url.Handle).First(&existingURL)
		log.Printf("Found existing handle: %v, to be replaced by %v", handle, url.Handle)
		if handle.Error == nil {
			if err := database.DB.Model(&existingURL).Updates(url).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("Error updating URL: %v", err)
				return
			}
		} else if errors.Is(handle.Error, gorm.ErrRecordNotFound) {
			log.Printf("Handle %v not found, inserting a new one", url.Handle)
			if err := database.DB.Create(&url).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("Error creating URL: %v", err)
				return
			}
		} else {
			http.Error(w, handle.Error.Error(), http.StatusInternalServerError)
			log.Printf("Unexpected error when trying to find the handle: %v", handle.Error)
			return
		}
	} else { // action is POST
		err := database.DB.Create(&url).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error creating URL: %v", err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "shortlink was stored successfully")
}
