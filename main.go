package main

import (
	"fmt"
	"net/http"
)

var sessions = make(map[string]string)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		sessionID := "some_cookie"

		sessions[sessionID] = "user_1"

		http.SetCookie(w, &http.Cookie{
			Name:  "session_id",
			Value: sessionID,
			Path:  "/",
		})
	})

	mux.HandleFunc("/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		clients := getClients()

		clientID := r.URL.Query().Get("client_id")
		redirectURL := r.URL.Query().Get("redirect_uri")
		responseType := r.URL.Query().Get("response_type")

		if _, ok := clients[clientID]; !ok {
			http.Error(w, "client is not defined", http.StatusBadRequest)
			return
		}

		if clients[clientID] != redirectURL {
			http.Error(w, "redirect url is not matching", http.StatusBadRequest)
			return
		}

		if responseType != "code" {
			http.Error(w, "response type is not valid", http.StatusBadRequest)
			return
		}

		userID, err := getUserFromRequest(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	})

	fmt.Println("server runs on 8080")
	http.ListenAndServe(":8080", mux)
}

func getClients() map[string]string {
	return map[string]string{
		"web_client": "http://localhost:8081/callback",
	}
}

func getUserFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil || cookie.Value == "" {
		return "", fmt.Errorf("no session")
	}

	userID, ok := sessions[cookie.Value]
	if !ok {
		return "", fmt.Errorf("invalid session")
	}

	return userID, nil
}
