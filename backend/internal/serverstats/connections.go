package serverstats

import (
	"net"
	"net/http"
	"sync/atomic"
)

type ConnectionWatcher struct {
	n atomic.Int64
}

func (cw *ConnectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew:
		cw.n.Add(1)

	case http.StateHijacked, http.StateClosed:
		cw.n.Add(-1)

	}
}

func (cw *ConnectionWatcher) Load() int64 {
	return cw.n.Load()
}
