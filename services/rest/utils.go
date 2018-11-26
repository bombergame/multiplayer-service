package rest

import (
	"errors"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

const (
	BasicPath       = "/multiplayer"
	RoomsPath       = BasicPath + "/rooms"
	RoomIDVar       = "room_id"
	RoomIDVarFormat = "[0-9-a-z]+"
	RoomPath        = RoomsPath + "/{" + RoomIDVar + ":" + RoomIDVarFormat + "}"
)

func (srv *Service) readRoomID(r *http.Request) (uuid.UUID, error) {
	idStr := mux.Vars(r)[RoomIDVar]
	if idStr == consts.EmptyString {
		err := errors.New(RoomIDVar + " not mapped")
		return uuid.Nil, errs.NewServiceError(err)
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, errs.NewInvalidFormatError("wrong " + RoomIDVar)
	}

	return id, nil
}