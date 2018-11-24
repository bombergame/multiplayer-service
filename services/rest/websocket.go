package rest

import (
	"fmt"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/objects/player"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
)

func (srv *Service) joinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		err = errs.NewServiceError(err)
		srv.WriteErrorWithBody(w, err)
		return
	}

	cmdCh := make(player.CommandsChan, 10)

	p := player.NewPlayer(rand.Int63())
	p.SetCommandsChan(&cmdCh)

	srv.room.AddPlayer(p)

	for {
		_, b, err := conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err) {
			srv.room.DeletePlayer(p)
			return
		}

		fmt.Println(string(b))
	}
}
