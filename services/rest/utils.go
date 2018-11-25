package rest

import (
	"errors"
	"github.com/bombergame/common/errs"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

func (srv *Service) readRoomID(r *http.Request) (uuid.UUID, error) {
	idStr := mux.Vars(r)["room_id"]
	if idStr == "" {
		err := errors.New("path variable room_id not mapped")
		return uuid.Nil, errs.NewServiceError(err)
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, errs.NewInvalidFormatError("wrong room_id")
	}

	return id, nil
}
