package handlers

import (
	"net/http"
	"text/template"
)

func (a *App) Signin(w http.ResponseWriter, r *http.Request) {
	a.Logger.Info("Getting uri params from the request url")
	responseType := r.URL.Query().Get("response_type")
	redirectURI := r.URL.Query().Get("redirect_uri")
	clientID := r.URL.Query().Get("client_id")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	a.Logger.Debug("URI Params", "response_type", responseType, "redirect_uri", redirectURI, "client_id", clientID, "scope", scope, "state", state)

	if r.Method == "GET" {
		a.Logger.Info("Creating signing uri with encoding")
		signinURI := CreateURI("/signin", clientID, responseType, redirectURI, scope, state)

		a.Logger.Debug("Created signin uri", "signin_uri", signinURI)

		a.Logger.Info("Rendering signing template, user should fill the form")
		template, _ := template.ParseFiles("templates/signin.html")
		template.Execute(w, struct {
			SubmitURI string
		}{
			SubmitURI: signinURI,
		})

		return
	}

	if r.Method != "POST" {
		http.Error(w, "wrong method call", http.StatusBadRequest)
		return
	}

	a.Logger.Info("Reading form values username and password")
	username := r.FormValue("username")
	password := r.FormValue("password")

	if _, ok := a.Users[username]; !ok {
		http.Error(w, "wrong username", http.StatusUnauthorized)
		return
	}
	a.Logger.Info("Username found")

	if a.Users[username] != password {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}
	a.Logger.Info("Password matching")

	sessionID, err := newSessionID()
	if err != nil {
		http.Error(w, "could not create session", http.StatusInternalServerError)
		return
	}
	a.Logger.Info("Generated session id")

	a.Logger.Info("Recoreded session id to the system")
	a.Sessions[sessionID] = username

	a.Logger.Info("Creating session cookie with session id value")
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	a.Logger.Info("Session cookie is created, now moving back to authorize")
	l := CreateURI("/oauth/authorize", clientID, responseType, redirectURI, scope, state)
	http.Redirect(w, r, l, http.StatusFound)
}
