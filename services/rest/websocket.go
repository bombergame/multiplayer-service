package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
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

	srv.Logger().Info("auth id: ", authID)

	_, err = srv.components.RoomsManager.GetRoom(roomID)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	for {
		_ = err
	}
}

func (srv *Service) handleAuthRequest(conn *websocket.Conn, msg *WebSocketRequest) (int64, error) {
	var authReqData AuthRequestData

	atItf, ok := msg.Data["auth_token"]
	if !ok {
		return 0, errs.NewInvalidFormatError("no auth_token field")
	}

	authReqData.AuthToken, ok = atItf.(string)
	if !ok {
		return 0, errs.NewInvalidFormatError("wrong auth_token field")
	}

	agItf, ok := msg.Data["user_agent"]
	if !ok {
		return 0, errs.NewInvalidFormatError("no user_agent field")
	}

	authReqData.UserAgent, ok = agItf.(string)
	if !ok {
		return 0, errs.NewInvalidFormatError("wrong user_agent field")
	}

	if atItf == consts.EmptyString {
		return -1, nil
	}

	pInfo, err := srv.components.AuthManager.GetProfileInfo(
		authReqData.AuthToken, authReqData.UserAgent)
	if err != nil {
		srv.writeWebSockError(conn, err)
		return 0, err
	}

	return pInfo.ID, nil
}

func (srv *Service) closeConnectionWithError(conn *websocket.Conn, err error) {
	srv.writeWebSockError(conn, err)
	if err := conn.Close(); err != nil {
		srv.Logger().Error(err)
	}
}
