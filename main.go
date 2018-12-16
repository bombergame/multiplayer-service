package main

import (
	"github.com/bombergame/multiplayer-service/game/fields"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/physics"
)

func main() {
	f := fields.NewField(physics.GetSize2D(100, 100))

	p := make(map[int64]*players.Player, 0)
	var i int64
	for i = 0; i < 10; i++ {
		p[i] = players.NewPlayer(i)
	}

	f.PlaceObjects(p)

	//logger := logs.NewLogger()
	//
	//restSrv := rest.NewService(
	//	rest.Config{},
	//	rest.Components{
	//		RoomsManager: utils.NewRoomsManager(),
	//		Components: restful.Components{
	//			Logger:      logger,
	//			AuthManager: utils.NewAuthManager(),
	//		},
	//	},
	//)
	//
	//ch := make(chan os.Signal)
	//signal.Notify(ch, os.Interrupt)
	//
	//go func() {
	//	if err := restSrv.Run(); err != nil {
	//		logger.Fatal(err)
	//	}
	//}()
	//
	//<-ch
	//
	//if err := restSrv.Shutdown(); err != nil {
	//	logger.Fatal(err)
	//}
}
