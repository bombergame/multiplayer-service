package rest

import (
	"github.com/bombergame/common/rest"
	"github.com/bombergame/multiplayer-service/config"
	"github.com/bombergame/multiplayer-service/game/room"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type Service struct {
	rest.Service
	room     *room.Room
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
		room: room.NewRoom(uuid.NewV4(), 2),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	mx := mux.NewRouter()
	mx.HandleFunc("/room", srv.joinRoom)

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
