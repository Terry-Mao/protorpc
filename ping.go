package protorpc

import (
	"errors"
	"log"
	"time"
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

// Reconnect for ping rpc server and reconnect with it when it's crash.
func Reconnect(dst **Client, quit chan struct{}, network, address string) {
	var (
		retires int
		tmp     *Client
		err     error
		call    *Call
		ch      = make(chan *Call, 1)
		client  = *dst
	)
	for {
		select {
		case call = <-client.Go(pingServiceMethod, nil, nil, ch).Done:
			if call.Error != nil {
				if tmp, err = Dial(network, address); err == nil {
					retires = 0
					*dst = tmp
					client = tmp
					log.Println("reconnect ok")
				} else {
					log.Println("reconnect retry")
					retires++
				}
			} else {
				// ping ok, reset retires
				log.Println("reconnect reset ok")
				retires = 0
			}
			time.Sleep(backoff(retires))
		case <-quit:
			return
		}
	}
}
