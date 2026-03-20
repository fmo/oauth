package main

import "net/http"

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	// get clients
	clients := getClients()

	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")

	if _, ok := clients[clientID]; !ok {
		http.Error(w, "client is not defined", http.StatusBadRequest)
		return
	}

	if clients[clientID] != redirectURI {
		http.Error(w, "redirect url is not matching", http.StatusBadRequest)
		return
	}

	if responseType != "code" {
		http.Error(w, "response type is not valid", http.StatusBadRequest)
		return
	}

	// get user
	userID, err := getUserFromRequest(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	code, _ := generateCode()
	storeCode(code, userID, clientID, redirectURI)

	redirect := redirectURI + "?code=" + code
	http.Redirect(w, r, redirect, http.StatusFound)
}
