package main

import (
	"github.com/bombergame/common/logs"
	"github.com/bombergame/common/registry"
	restful "github.com/bombergame/common/rest"
	"github.com/bombergame/multiplayer-service/config"
	"github.com/bombergame/multiplayer-service/services/rest"
	"github.com/bombergame/multiplayer-service/utils"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	regClient := registry.NewClient(
		registry.Config{
			RegistryAddress: config.RegistryAddress,
			ServiceName:     config.ServiceName,
			ServicePort:     config.HttpPort,
			ServiceHost: func() string {
				host, err := os.Hostname()
				if err != nil {
					panic(err)
				}
				return host
			}(),
		},
	)
	if err := regClient.Register(); err != nil {
		logger.Fatal(err)
		return
	}
	defer func() {
		if err := regClient.Deregister(); err != nil {
			logger.Fatal(err)
		}
	}()

	restSrv := rest.NewService(
		rest.Config{},
		rest.Components{
			RoomsManager: utils.NewRoomsManager(),
			Components: restful.Components{
				Logger:      logger,
				AuthManager: utils.NewAuthManager(),
			},
		},
	)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	go func() {
		if err := restSrv.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-ch

	if err := restSrv.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
