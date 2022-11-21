package main

import (
	"os"

	"github.com/4ynyky/router/internal/config"
	radiustester "github.com/4ynyky/router/pkg/servers/radiusTester"
	radiusservice "github.com/4ynyky/router/pkg/service/radiusService"
	log "github.com/sirupsen/logrus"
)

func main() {
	argsWithoutProg := os.Args[1:]

	configFileName := ""
	for i, v := range argsWithoutProg {
		if v == "-f" {
			if (len(argsWithoutProg) - 1) > i {
				configFileName = argsWithoutProg[i+1]
			} else {
				log.Error("Not found config file name after -f")
				os.Exit(1)
			}
		}
	}
	if len(configFileName) == 0 {
		log.Error("Not found param -f")
		os.Exit(1)
	}

	config := config.LoadConfiguration(configFileName)

	radClientConfig := &radiustester.RadiusClientConfig{}
	radClientConfig.Secret = config.Secret
	radClientConfig.Host = config.Host + ":" + config.AuthPort

	for _, v := range config.Auth {
		radClientConfig.Username = v.Username
		radClientConfig.Password = v.Password

		radService := radiusservice.NewRadiusService(radiustester.NewRadiusTransport(radClientConfig))
		radService.TestAuth()
	}

	// radClientConfig.AcctHost = "localhost:1813"
	// radClientConfig.AcctSessionID = "GGG"

	//radService.TestAcct()
}
