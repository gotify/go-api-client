package auth

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime"
)

func BasicAuth(username, password string) runtime.ClientAuthInfoWriter  {
	return httptransport.BasicAuth(username, password)
}

func TokenAuth(token string) runtime.ClientAuthInfoWriter {
	return httptransport.APIKeyAuth("X-Gotify-Key", "header", token)
}