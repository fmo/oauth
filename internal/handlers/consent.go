package handlers

import (
	"net/http"
	"text/template"
)

func (a *App) Consent(w http.ResponseWriter, r *http.Request) {
	scope := r.URL.Query().Get("scope")

	template, err := template.ParseFiles("templates/consent.html")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	template.Execute(w, scope)
}
