package rest

import (
	"github.com/bombergame/multiplayer-service/domains"
	"github.com/bombergame/multiplayer-service/game/rooms"
	"net/http"
)

func (srv *Service) createRoom(w http.ResponseWriter, r *http.Request) {
	_, err := srv.ReadAuthProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	var mRoom domains.Room
	if err := srv.ReadRequestBody(&mRoom, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}
	if err := mRoom.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	room := rooms.NewRoom(mRoom)
	n, err := srv.components.RoomsManager.AddRoom(room)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}
	room.RunGame()

	srv.metrics.numRooms.Set(float64(n))

	mRoom.ID = room.ID()
	srv.WriteOkWithBody(w, mRoom)
}
