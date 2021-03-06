package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/rest"
	"github.com/bombergame/multiplayer-service/config"
	"github.com/bombergame/multiplayer-service/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Service struct {
	rest.Service
	config     Config
	components Components
	metrics    Metrics
	upgrader   websocket.Upgrader
}

type Config struct {
	rest.Config
}

type Components struct {
	rest.Components
	RoomsManager *utils.RoomsManager
}

type Metrics struct {
	numRooms prometheus.Gauge
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
		metrics: Metrics{
			numRooms: NewNumRooms(),
		},
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	mx := mux.NewRouter()
	mx.Handle("/multiplayer/rooms", handlers.MethodHandler{
		http.MethodPost: srv.WithAuth(http.HandlerFunc(srv.createRoom)),
	})
	mx.Handle("/multiplayer/rooms/{room_id:[0-9-a-z]+}/ws", http.HandlerFunc(srv.handleGameplay))
	mx.Handle("/metrics", promhttp.Handler())

	srv.SetHandler(srv.WithLogs(srv.WithRecover(mx)))

	return srv
}
