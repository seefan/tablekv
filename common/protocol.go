package common

type NetLayout interface {
	Start(pm GetProcessor, host string, port int,timeout int) error
	Stop() error
}
