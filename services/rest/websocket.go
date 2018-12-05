package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func (srv *Service) handleGameplay(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.Logger().Info("connection upgraded")

	roomID, err := srv.readRoomID(r)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.Logger().Info("room id: ", roomID)

	_, b, err := conn.ReadMessage()
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.Logger().Info("request message: ", string(b))

	var msg WebSocketRequest
	if err := easyjson.Unmarshal(b, &msg); err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.Logger().Info("message: ", msg)

	if msg.Type != "auth" {
		err := errs.NewNotAuthorizedError()
		srv.closeConnectionWithError(conn, err)
		return
	}

	authID, err := srv.handleAuthRequest(conn, &msg)
	if err != nil {
		return
	}

	room, err := srv.components.RoomsManager.GetRoom(roomID)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	p := players.NewPlayer(authID)

	if err := room.AddPlayer(p); err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.writeWebSockOk(conn)
}

func (srv *Service) handleAuthRequest(conn *websocket.Conn, msg *WebSocketRequest) (int64, error) {
	var authReqData AuthRequestData

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: &authReqData, TagName: "json",
	})
	if err != nil {
		return consts.AnyInt, errs.NewServiceError(err)
	}

	if err := decoder.Decode(&msg.Data); err != nil {
		return consts.AnyInt, errs.NewInvalidFormatError("wrong auth message")
	}

	if authReqData.AuthToken == consts.EmptyString {
		return -1, nil
	}

	pInfo, err := srv.components.AuthManager.GetProfileInfo(
		authReqData.AuthToken, authReqData.UserAgent)
	if err != nil {
		return consts.AnyInt, err
	}

	return pInfo.ID, nil
}

func (srv *Service) closeConnectionWithError(conn *websocket.Conn, err error) {
	srv.writeWebSockError(conn, err)
	if err := conn.Close(); err != nil {
		srv.Logger().Error(err)
	}
}
