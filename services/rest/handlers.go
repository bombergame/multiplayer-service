package rest

import (
	"github.com/bombergame/multiplayer-service/game/room"
	"github.com/satori/go.uuid"
	"net/http"
)

func (srv *Service) createRoom(w http.ResponseWriter, r *http.Request) {
	_, err := srv.ReadAuthProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	var newRoom Room
	if err := srv.ReadRequestBody(&newRoom, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}
	if err := newRoom.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	newRoom.ID = uuid.NewV4()

	err = srv.components.RoomsManager.AddRoom(room.NewRoom(newRoom.ID))
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, newRoom)
}

func (srv *Service) getRooms(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) getRoom(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) joinRoom(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func (srv *Service) deleteRoom(w http.ResponseWriter, r *http.Request) {
	//TODO
}