package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func (a *App) ListSessions(w http.ResponseWriter, r *http.Request) {
	for session, user := range a.Sessions {
		fmt.Printf("Session: %s, User: %s\n", session, user)
	}
}

func newSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
