package rest

import (
	"github.com/bombergame/common/consts"
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
	cf.Host, cf.Port = consts.EmptyString, config.HttpPort

	srv := &Service{
		Service: *rest.NewService(
			cf.Config,
			cpn.Components,
		),
		config:     cf,
		components: cpn,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
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

	cors := rest.CORS{
		Origins:     []string{"http://127.0.0.1:8000"},
		Methods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		Headers:     []string{"User-Agent", "Authorization", "Content-Type", "Content-Length"},
		Credentials: true,
	}

	srv.SetHandler(srv.WithLogs(srv.WithCORS(srv.WithRecover(mx), cors)))

	return srv
}
