package handlers

import (
	"net/http"
	"text/template"

	"github.com/fmo/oauth/internal"
)

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		responseType := r.URL.Query().Get("response_type")
		redirectURI := r.URL.Query().Get("redirect_uri")
		clientID := r.URL.Query().Get("client_id")
		scope := r.URL.Query().Get("scope")

		loginURI := internal.CreateLoginURI(clientID, responseType, redirectURI, scope)

		template, _ := template.ParseFiles("templates/login.html")

		template.Execute(w, struct {
			SubmitURI string
		}{
			SubmitURI: loginURI,
		})

		return
	}

	if r.Method != "POST" {
		http.Error(w, "wrong method call", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	redirectURI := r.URL.Query().Get("redirect_uri")

	if _, ok := a.Users[username]; !ok {
		http.Error(w, "wrong username", http.StatusUnauthorized)
		return
	}

	if a.Users[username] != password {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	sessionID, err := newSessionID()
	if err != nil {
		http.Error(w, "could not create session", http.StatusInternalServerError)
		return
	}

	a.Sessions[sessionID] = username

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, redirectURI, http.StatusFound)
}
