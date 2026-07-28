package handlers

import (
    "net/http"
    "html/template"
    "log"
    "path/filepath"
)

func GetURLUpdates (path string) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        if !IsValidMethod(w, r, http.MethodGet) {
            return
        }

        tmplPath := filepath.Join(path, "update_form.html")
        tmpl, err := template.ParseFiles(tmplPath)
        if err != nil {
            http.Error(w, "Could not load template", http.StatusInternalServerError)
            log.Printf("Template error: %v", err)
            return
        }

        // Render the template to the response
        err = tmpl.Execute(w, nil)
        if err != nil {
            http.Error(w, "Could not render template", http.StatusInternalServerError)
            log.Printf("Render error: %v", err)
        }
    }
}
