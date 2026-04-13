## History

2007 first OAuth 1.0 released

## Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant Browser
    participant AuthServer

    Client->>Browser: Open /oauth/authorize\n(client_id, redirect_uri, scope, state)
    Browser->>AuthServer: GET /oauth/authorize

    alt No session cookie
        AuthServer-->>Browser: Redirect to /login

        Browser->>AuthServer: POST /login (username, password)
        AuthServer-->>Browser: Set session cookie + redirect back to /authorize
    end

    Browser->>AuthServer: GET /oauth/authorize (again with session)

    alt User consent required
        AuthServer-->>Browser: Show consent screen
        Browser->>AuthServer: Submit consent (approve)
    end

    AuthServer-->>Browser: Redirect to redirect_uri?code=XYZ&state=ABC
    Browser-->>Client: Return with authorization code
```

## Use cases

One of them is single sign on for multiple apps

## Example Client

An example client for this auth server is here:

https://github.com/fmo/oauth2_client

## Request Params for /oauth/authorize

### response_type 

response_type usually is `code` which is a code being returned in case of successful requests. Later that will be exchanged with a token.

## Authorization Server

    * Authenticates users(login)
    * Issues Authorization codes
    * Exchanges them for tokens
    * Validate clients
    * Returns access/id/refresh tokens

## Resource Server

Your APIs

## OAuth vs OpenID Connect

OpenID Connect is a extension to OAuth.

You call first `/oauth/authorize` then you will get a `code` but if you indicate in the call you will need an openid token then:

```
GET /authorize?
  response_type=code
  &client_id=...
  &scope=openid email profile
```

Then you will be able to get from `/oauth/token` id token as well.

```
{
  "access_token": "xyz123",
  "id_token": "jwt_token_here",
  "token_type": "Bearer",
  "expires_in": 3600
}
```

## More Terminology

* Resource Owner: The user
* User Agent: Device
* OAuth Client: The applicaition
* Resource Server: The API

## Grant Types

* client_credentials: M2M communication
* authorization_code: User Login Flow (OAuth2/OIDC)
