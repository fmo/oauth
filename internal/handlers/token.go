package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (a *App) Token(w http.ResponseWriter, r *http.Request) {
	clientID := r.FormValue("client_id")
	clientSecret := r.FormValue("client_secret")
	redirectURI := r.FormValue("redirect_uri")
	code := r.FormValue("code")

	if _, ok := a.Clients[clientID]; !ok {
		http.Error(w, "client does not exist", http.StatusBadRequest)
		return
	}

	client := a.Clients[clientID]

	if client.Secret != clientSecret {
		http.Error(w, "wrong client secret", http.StatusUnauthorized)
		return
	}

	grantType := r.FormValue("grant_type")

	switch grantType {
	case "client_credentials":
		w.Write([]byte(`{
			"access_token": "abc123",
			"token_type": "Bearer",
			"expires_in": 3600
		}`))
	case "authorization_code":
		authCode, err := a.ConsumeCode(code, clientID, redirectURI)
		if err != nil {
			http.Error(w, "code is wrong", http.StatusUnauthorized)
			return
		}

		token, err := a.GenerateCode()
		if err != nil {
			http.Error(w, "cant generate token", http.StatusInternalServerError)
			return
		}

		a.StoreToken(token, authCode.UserID, authCode.ClientID, authCode.Scope)

		var isOIDC bool
		scopes := strings.Fields(authCode.Scope)
		for _, scope := range scopes {
			if scope == "openid" {
				isOIDC = true
			}
		}

		resp := map[string]any{
			"access_token": token,
			"token_type":   "Bearer",
			"expires_in":   3600,
		}

		if isOIDC {
			claims := jwt.MapClaims{
				"sub": authCode.UserID,
				"iss": "localhost:8080",
				"aud": clientID,
				"exp": time.Now().Add(time.Hour).Unix(),
				"iat": time.Now().Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			idToken, err := token.SignedString([]byte("my-secret"))
			if err != nil {
				http.Error(w, "cant create id token", http.StatusInternalServerError)
				return
			}

			resp["id_token"] = idToken
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	case "refresh_token":
	default:
		http.Error(w, "unsupported grant", http.StatusBadRequest)
		return
	}
}
