package gotify

import (
	"net/url"
	"net/http"
	api "github.com/gotify/go-api-client/v2/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func NewClient(url *url.URL, client *http.Client) *api.GotifyREST {
	runtime := httptransport.NewWithClient(url.Host, url.Path, []string{url.Scheme}, client)
	return api.New(runtime, nil)
}