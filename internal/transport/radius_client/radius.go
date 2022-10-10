package radius_client

import (
	"context"
	"log"

	"github.com/4ynyky/router/internal/transport"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

var (
	ClientUsername = "tim"
	ClientPassword = "12345"
)

type radiusTransport struct {
}

func NewRadiusTransport() transport.ITransport {
	return &radiusTransport{}
}

func (rt *radiusTransport) SendAuth() error {
	packet := radius.New(radius.CodeAccessRequest, []byte(`secret`))
	rfc2865.UserName_SetString(packet, ClientUsername)
	rfc2865.UserPassword_SetString(packet, ClientPassword)
	response, err := radius.Exchange(context.Background(), packet, "localhost:1812")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Code:", response.Code)

	return nil
}
