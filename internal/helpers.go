package internal

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

func CreateLoginURI(clientID, responseType, redirectURI, scope string) string {
	u, _ := url.Parse("/login")
	q := u.Query()
	q.Set("client_id", clientID)
	q.Set("response_type", responseType)
	q.Set("redirect_uri", redirectURI)
	q.Set("scope", scope)

	u.RawQuery = q.Encode()

	return u.String()
}
