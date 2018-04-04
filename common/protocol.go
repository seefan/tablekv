package common

type NetLayout interface {
	Start(pm GetProcessor, host string, port int) error
	Stop() error
}
