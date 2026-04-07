package handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetUserFromRequest(r *http.Request, sessions map[string]string) (string, error) {
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

func CreateURI(base, clientID, responseType, redirectURI, scope, state string) string {
	u, _ := url.Parse(base)

	q := u.Query()
	q.Add("client_id", clientID)
	q.Add("response_type", responseType)
	q.Add("redirect_uri", redirectURI)
	q.Add("scope", scope)
	q.Add("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}

func CreateRedirectURI(redirectURI, code, state string) string {
	u, _ := url.Parse(redirectURI)

	q := u.Query()
	q.Add("code", code)
	q.Add("state", state)

	u.RawQuery = q.Encode()

	return u.String()
}
