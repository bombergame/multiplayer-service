package rest

import (
	"context"
	"net/http"

	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
)

//Service is http server wrapper
type Service struct {
	config     Config
	components Components
	server     http.Server
}

//Config contains service configuration parameters
type Config struct {
	Host string
	Port string
}

//Components contains service components
type Components struct {
	Logger      *logs.Logger
	AuthManager auth.AuthenticationManager
}

//NewService creates service instance
func NewService(config Config, components Components) *Service {
	return &Service{
		config:     config,
		components: components,
		server: http.Server{
			Addr: config.Host + ":" + config.Port,
		},
	}
}

//Run starts the service
func (srv *Service) Run() error {
	srv.components.Logger.Info("http service running on: " + srv.server.Addr)
	return srv.server.ListenAndServe()
}

//Shutdown forces the service to stop
func (srv *Service) Shutdown() error {
	srv.components.Logger.Info("http service shutdown initialized")
	return srv.server.Shutdown(context.TODO())
}

//SetHandler sets the http connections handler
func (srv *Service) SetHandler(h http.Handler) {
	srv.server.Handler = h
}

//Logger returns the service logger
func (srv *Service) Logger() *logs.Logger {
	return srv.components.Logger
}
