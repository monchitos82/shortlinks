package handlers

import (
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {

	if !IsValidMethod(w, r, http.MethodGet) {
		return
	}
	w.WriteHeader(http.StatusOK)
}