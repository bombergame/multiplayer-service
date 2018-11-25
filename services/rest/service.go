package rest

import (
	"github.com/bombergame/common/rest"
	"github.com/bombergame/multiplayer-service/config"
	"github.com/bombergame/multiplayer-service/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

type Service struct {
	rest.Service
	config     Config
	components Components
	upgrader   websocket.Upgrader
}

type Config struct {
	rest.Config
}

type Components struct {
	rest.Components
	RoomsManager *utils.RoomsManager
}

func NewService(cf Config, cpn Components) *Service {
	cf.Host, cf.Port = "", config.HttpPort

	srv := &Service{
		Service: *rest.NewService(
			cf.Config,
			cpn.Components,
		),
		config:     cf,
		components: cpn,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	mx := mux.NewRouter()
	mx.Handle(RoomsPath, handlers.MethodHandler{
		http.MethodPost: srv.WithAuth(http.HandlerFunc(srv.createRoom)),
	})
	mx.Handle(RoomPath, handlers.MethodHandler{
		http.MethodDelete: srv.WithAuth(http.HandlerFunc(srv.deleteRoom)),
	})
	mx.Handle(RoomPath+"/ws", http.HandlerFunc(srv.handleGameplay))

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
