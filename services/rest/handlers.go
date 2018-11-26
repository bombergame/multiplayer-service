package rest

import (
	"github.com/bombergame/multiplayer-service/domains"
	"github.com/bombergame/multiplayer-service/game/rooms"
	"github.com/satori/go.uuid"
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

	mRoom.ID = uuid.NewV4()
	room := rooms.NewRoom(mRoom)

	err = srv.components.RoomsManager.AddRoom(room)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, mRoom)
}

func (srv *Service) deleteRoom(w http.ResponseWriter, r *http.Request) {
	//TODO
}
