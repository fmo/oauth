// Package handlers take care of requests
package handlers

import (
	"net/http"
)

func (a *App) Authorize(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	redirectURI := r.URL.Query().Get("redirect_uri")
	responseType := r.URL.Query().Get("response_type")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	// validate
	if _, ok := a.Clients[clientID]; !ok {
		http.Error(w, "client is not defined", http.StatusBadRequest)
		return
	}
	if a.Clients[clientID].RedirectURI != redirectURI {
		http.Error(w, "redirect url is not matching", http.StatusBadRequest)
		return
	}
	if responseType != "code" {
		http.Error(w, "response type is not valid", http.StatusBadRequest)
		return
	}

	// get user
	userID, err := GetUserFromRequest(r, a.Sessions)
	if err != nil {
		loginURI := CreateURI("/login", clientID, responseType, redirectURI, scope, state)
		http.Redirect(w, r, loginURI, http.StatusFound)
		return
	}

	// generate code
	code, err := a.GenerateCode()
	if err != nil {
		http.Error(w, "cant generate code", http.StatusInternalServerError)
		return
	}
	a.StoreCode(code, userID, clientID, redirectURI, scope)

	// create redirect uri
	rduri := CreateRedirectURI(redirectURI, code, state)

	// return back to client
	http.Redirect(w, r, rduri, http.StatusFound)
}
