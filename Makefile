authorize:
	curl -X GET "http://localhost:8080/oauth/authorize?response_type=code&redirect_uri=http://localhost:8081/callback&client_id=web_client"

m2m:
	curl -X POST http://localhost:8080/oauth/token \
  	-H "Content-Type: application/x-www-form-urlencoded" \
  	-d "grant_type=client_credentials" \
  	-d "client_id=web_client" \
  	-d "client_secret=axaa"
