package main

import (
	"github.com/bombergame/common/logs"
	"github.com/bombergame/multiplayer-service/services/rest"
	"github.com/bombergame/multiplayer-service/utils"
	"os"
	"os/signal"
)

func main() {
	logger := logs.NewLogger()

	restSrv := rest.NewService(
		rest.Config{},
		rest.Components{
			RoomsManager: utils.NewRoomsManager(),
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
