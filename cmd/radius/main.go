package main

import (
	"github.com/4ynyky/router/internal/transport/radius_client"
)

func main() {
	radius_client.NewRadiusTransport()
	radius_client.NewRadiusTransport().SendAuth()
}
