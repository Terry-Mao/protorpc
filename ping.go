package protorpc

import (
	"errors"
)

const (
	pingServiceMethod = "protorpc.Ping"
	pingService       = "protorpc"
)

var (
	DefaultPinger = &Pinger{}
	ErrProtoRPC   = errors.New(pingService + " is a inner service")
)

type Pinger struct {
}

func (p *Pinger) Ping(args *Ping, reply *Ping) error {
	return nil
}
