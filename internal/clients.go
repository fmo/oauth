package internal

type Client struct {
	Secret      string
	RedirectURI string
}

func GetClients() map[string]Client {
	return map[string]Client{
		"web_client": Client{
			Secret:      "axaa",
			RedirectURI: "http://localhost:8081/callback",
		},
	}
}
