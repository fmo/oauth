m2m:
	curl -X POST http://localhost:8080/oauth/token \
  	-H "Content-Type: application/x-www-form-urlencoded" \
  	-d "grant_type=client_credentials" \
  	-d "client_id=web_client" \
  	-d "client_secret=secret"
