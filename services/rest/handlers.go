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

	var room domains.Room
	if err := srv.ReadRequestBody(&room, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}
	if err := room.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	room.ID = uuid.NewV4()

	err = srv.components.RoomsManager.AddRoom(rooms.NewRoom(room.ID))
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, room)
}

func (srv *Service) deleteRoom(w http.ResponseWriter, r *http.Request) {
	//TODO
}
