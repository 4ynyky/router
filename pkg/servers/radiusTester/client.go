package radiustester

import (
	"context"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
	"layeh.com/radius/rfc2866"
)

type IRadiusClient interface {
	SendAuth() (*radius.Packet, error)
	SendAcctStart() (*radius.Packet, error)
	SendAcctInterimUpdate() (*radius.Packet, error)
}

type RadiusClientConfig struct {
	Host     string
	Secret   string
	Username string
	Password string

	AcctHost      string
	AcctSessionID string
}

type radiusClient struct {
	rcc *RadiusClientConfig
}

func NewRadiusTransport(rcc *RadiusClientConfig) IRadiusClient {
	return &radiusClient{rcc: rcc}
}

func (rt *radiusClient) setAuthData(packet *radius.Packet) {
	rfc2865.UserName_SetString(packet, rt.rcc.Username)
	rfc2865.UserPassword_SetString(packet, rt.rcc.Password)
}

func (rt *radiusClient) SendAuth() (*radius.Packet, error) {
	packet := radius.New(radius.CodeAccessRequest, []byte(rt.rcc.Secret))
	rt.setAuthData(packet)
	response, err := radius.Exchange(context.Background(), packet, rt.rcc.Host)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (rt *radiusClient) SendAcctStart() (*radius.Packet, error) {
	packet := radius.New(radius.CodeAccountingRequest, []byte(rt.rcc.Secret))
	rt.setAuthData(packet)

	rfc2866.AcctSessionID_Add(packet, []byte(rt.rcc.AcctSessionID))
	rfc2866.AcctStatusType_Add(packet, rfc2866.AcctStatusType_Value_Start)
	response, err := radius.Exchange(context.Background(), packet, rt.rcc.AcctHost)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (rt *radiusClient) SendAcctInterimUpdate() (*radius.Packet, error) {
	packet := radius.New(radius.CodeAccountingRequest, []byte(rt.rcc.Secret))
	rt.setAuthData(packet)

	rfc2866.AcctSessionID_Add(packet, []byte(rt.rcc.AcctSessionID))
	rfc2866.AcctStatusType_Add(packet, rfc2866.AcctStatusType_Value_InterimUpdate)
	response, err := radius.Exchange(context.Background(), packet, rt.rcc.AcctHost)
	if err != nil {
		return nil, err
	}
	return response, nil
}
