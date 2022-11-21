package radiusservice

import (
	log "github.com/sirupsen/logrus"
)

func (rs *radiusService) TestAuth() {
	log.Println("Test auth")
	res, err := rs.radClient.SendAuth()
	if err != nil {
		log.Error("Failed rs.radClient.SendAuth: ", err.Error())
		return
	}

	logRadPacket(res)
}
