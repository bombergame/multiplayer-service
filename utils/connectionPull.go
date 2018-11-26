package utils

import (
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
	"sync"
)

type ConnectionPull struct {
	mu    *sync.RWMutex
	conns map[int64]*websocket.Conn
}

func NewConnectionPull() *ConnectionPull {
	return &ConnectionPull{
		mu:    &sync.RWMutex{},
		conns: make(map[int64]*websocket.Conn),
	}
}

func (p *ConnectionPull) AddConnection(id int64, conn *websocket.Conn) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.conns[id]; ok {
		_ = p.conns[id].Close()
	}
	p.conns[id] = conn

	return nil
}

func (p *ConnectionPull) ListenAndServe() {

}

func (p *ConnectionPull) Broadcast(v easyjson.Marshaler) {

}
