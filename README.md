# Go API client for swagger [![travis-badge]][travis]

Generated with [swagger-go](https://github.com/go-swagger/go-swagger).

## Example

```go
package main

import (
	"github.com/gotify/go-api-client/gotify"
	"net/url"
	"net/http"
	"log"
	"github.com/gotify/go-api-client/client/message"
	"github.com/gotify/server/model"
	"github.com/gotify/go-api-client/auth"
)

const (
	gotifyURL ="https://push.gotify.url"
	applicationToken = "AxNVvRfx9.ZKCTj"
)

func main() {
	myURL, _ := url.Parse(gotifyURL)
	client := gotify.NewClient(myURL, &http.Client{})
	versionResponse, err := client.Version.GetVersion(nil)

	if err != nil {
		log.Fatal("Could not request version ", err)
	}
	version := versionResponse.Payload
	log.Println("Found version", version.Version)

	params := message.NewCreateMessageParams()
	params.Body = &model.Message{
		Title: "my title",
		Message: "my message",
		Priority: 5,
	}
	_, err = client.Message.CreateMessage(params, auth.TokenAuth(applicationToken))

	if err != nil {
		log.Fatal("Could not send message ", err)
	}
	log.Println("Message Sent!")
}
```

## Update Instructions

* Change version in [Makefile](Makefile).
* Run `make clean generate`
* Commit changes

 [travis]: https://travis-ci.org/gotify/go-api-client
 [travis-badge]: https://travis-ci.org/gotify/go-api-client.svg?branch=master