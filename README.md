# Gotify Golang API-Client [![travis-badge]][travis]

Generated with [swagger-go](https://github.com/go-swagger/go-swagger).

## Example

```go
package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

const (
	gotifyURL        = "https://push.gotify.url"
	applicationToken = "AxNVvRfx9.ZKCTj"
)

func main() {
	myURL, _ := url.Parse(gotifyURL)
	client := gotify.NewClient(myURL, &http.Client{})
	versionResponse, err := client.Version.GetVersion(nil)

	if err != nil {
		log.Fatal("Could not request version ", err)
		return
	}
	version := versionResponse.Payload
	log.Println("Found version", *version)

	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:    "my title",
		Message:  "my message",
		Priority: 5,
	}
	_, err = client.Message.CreateMessage(params, auth.TokenAuth(applicationToken))

	if err != nil {
		log.Fatalf("Could not send message %v", err)
		return
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
