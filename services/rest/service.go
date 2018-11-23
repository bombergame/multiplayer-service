package rest

import (
	"github.com/bombergame/common/rest"
	"github.com/bombergame/multiplayer-service/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Service struct {
	rest.Service
	upgrader websocket.Upgrader
}

type Config struct {
	rest.Config
}

type Components struct {
	rest.Components
}

func NewService(cf Config, cpn Components) *Service {
	cf.Host, cf.Port = "", config.HttpPort

	srv := &Service{
		Service: *rest.NewService(
			cf.Config,
			cpn.Components,
		),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	mx := mux.NewRouter()
	mx.HandleFunc("/rooms/{room_id::[0-9]+}", srv.joinRoom)

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
