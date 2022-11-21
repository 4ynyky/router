package radiusservice

import (
	radiustester "github.com/4ynyky/router/pkg/servers/radiusTester"
	log "github.com/sirupsen/logrus"
	"layeh.com/radius"
)

type IRadiusService interface {
	TestAuth()
	TestAcct()
}

type radiusService struct {
	radClient radiustester.IRadiusClient
}

func NewRadiusService(radClient radiustester.IRadiusClient) IRadiusService {
	return &radiusService{radClient: radClient}
}

func logRadPacket(packet *radius.Packet) {
	log.Println("Code: ", packet.Code)
	for _, v := range packet.Attributes {
		log.Println("	Type: ", v.Type)
		log.Println("		Value: ", v.Attribute)
	}
}
