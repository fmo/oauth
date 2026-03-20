package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

type AuthCode struct {
	UserID      string
	ClientID    string
	RedirectURI string
	ExpiresAt   time.Time
}

var codeStore = map[string]AuthCode{}

func generateCode() (string, error) {
	b := make([]byte, 32) // 256-bit

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// URL-safe (no + / =)
	code := base64.RawURLEncoding.EncodeToString(b)

	return code, nil
}

func storeCode(code, userID, clientID, redirectURI string) {
	codeStore[code] = AuthCode{
		UserID:      userID,
		ClientID:    clientID,
		RedirectURI: redirectURI,
		ExpiresAt:   time.Now().Add(2 * time.Minute),
	}
}

func consumeCode(code, clientID, redirectURI string) (*AuthCode, error) {
	data, ok := codeStore[code]
	if !ok {
		return nil, fmt.Errorf("invalid code")
	}

	// check expiration
	if time.Now().After(data.ExpiresAt) {
		delete(codeStore, code)
		return nil, fmt.Errorf("code expired")
	}

	// check client binding
	if data.ClientID != clientID {
		return nil, fmt.Errorf("client mismatch")
	}

	// check redirect URI (important!)
	if data.RedirectURI != redirectURI {
		return nil, fmt.Errorf("redirect mismatch")
	}

	// one-time use → delete immediately
	delete(codeStore, code)

	return &data, nil
}
