package transport

type ITransport interface {
	SendAuth() error
}
