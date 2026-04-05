package internal

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type AuthCode struct {
	UserID      string
	ClientID    string
	RedirectURI string
	ExpiresAt   time.Time
	Scope       string
}

func GenerateCode() (string, error) {
	b := make([]byte, 32) // 256-bit

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// URL-safe (no + / =)
	code := base64.RawURLEncoding.EncodeToString(b)

	return code, nil
}
