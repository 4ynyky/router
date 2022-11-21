package radiusservice

import log "github.com/sirupsen/logrus"

func (rs *radiusService) TestAcct() {
	log.Println("Test acct")
	res, err := rs.radClient.SendAcctStart()
	if err != nil {
		log.Error("Failed rs.radClient.SendAcctStart: ", err.Error())
		return
	}

	logRadPacket(res)
}
