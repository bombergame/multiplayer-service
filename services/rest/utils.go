package rest

import (
	"errors"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
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
		return uuid.Nil, errs.NewInternalServiceError(err)
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, errs.NewInvalidFormatError("wrong " + RoomIDVar)
	}

	return id, nil
}

func (srv *Service) writeWebSockOk(conn *websocket.Conn) {
	srv.writeWebSockJSON(conn, &ws.OutMessage{
		Type: ws.OkMessageType,
		Data: ws.OkMessageData{
			Message: "ok",
		},
	})
}

func (srv *Service) writeWebSockError(conn *websocket.Conn, err error) {
	srv.writeWebSockJSON(conn, &ws.OutMessage{
		Type: ws.ErrorMessageType,
		Data: ws.ErrorMessageData{
			Message: err.Error(),
		},
	})
}

func (srv *Service) writeWebSockJSON(conn *websocket.Conn, v easyjson.Marshaler) {
	resp, err := easyjson.Marshal(v)
	if err != nil {
		panic(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, resp); err != nil {
		srv.Logger().Error(err)
		return
	}
}
