package handlers

import (
	"encoding/base64"
	"go-shortlinks/database"
	"go-shortlinks/models"
	"log"
	"net/http"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if !IsValidMethod(w, r, http.MethodGet) {
		return
	}
	handle := r.URL.Path[1:] // Remove leading "/"

	var url models.URL
	err := database.DB.Where("handle = ?", handle).First(&url).Error
	if err != nil {
		http.NotFound(w, r)
		log.Printf("Error finding URL: %v", err)
		return
	}
	decodedURL, err := base64.StdEncoding.DecodeString(url.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error decoding URL: %v", err)
		return
	}
	http.Redirect(w, r, string(decodedURL), http.StatusFound)
}
