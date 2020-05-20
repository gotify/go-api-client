package gotify

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	api "github.com/gotify/go-api-client/v2/client"
)

var GOTIFY_SKIP_TLS = "GOTIFY_SKIP_TLS"

func NewClient(url *url.URL, client *http.Client) *api.GotifyREST {
	should_skip_tls := os.Getenv(GOTIFY_SKIP_TLS) == "True"

	if should_skip_tls {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	runtime := httptransport.NewWithClient(url.Host, url.Path, []string{url.Scheme}, client)
	return api.New(runtime, nil)
}
