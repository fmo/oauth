package main

import (
	"net/http"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	grantType := r.FormValue("grant_type")

	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")

	clients := getClients()

	switch grantType {
	case "client_credentials":
		if v, ok := clients[clientID]; ok {
			if v.Secret != clientSecret {
				http.Error(w, "not matching secret", http.StatusUnauthorized)
				return
			}
		}
	case "authorization_code":
	case "refresh_token":
	default:
		http.Error(w, "unsupported grant", http.StatusBadRequest)
		return
	}
}
