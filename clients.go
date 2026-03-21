package main

type Client struct {
	Secret      string
	RedirectURI string
}

func getClients() map[string]Client {
	return map[string]Client{
		"web_client": Client{
			Secret:      "axaa",
			RedirectURI: "http://localhost:8081/callback",
		},
	}
}
