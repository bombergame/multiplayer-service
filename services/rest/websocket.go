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
		srv.WriteErrorWithBody(w, errs.NewServiceError(err))
		return
	}

	roomID, err := srv.readRoomID(r)
	if err != nil {
		srv.writeWebSockError(conn, err)
		return
	}

	_, b, err := conn.ReadMessage()
	if err != nil {
		srv.writeWebSockError(conn, err)
		return
	}

	var msg WebSocketMessage
	if err := easyjson.Unmarshal(b, &msg); err != nil {
		srv.writeWebSockError(conn, err)
		return
	}

	_, err = srv.components.RoomsManager.GetRoom(roomID)
	if err != nil {
		srv.writeWebSockError(conn, err)
		return
	}

	for {
		_ = err
	}
}

func (srv *Service) handleAuthRequest(conn *websocket.Conn, msg *WebSocketMessage) (int64, error) {
	var authReqData AuthRequestData
	if err := easyjson.Unmarshal(msg.Data, &authReqData); err != nil {
		srv.writeWebSockError(conn, err)
		return 0, err
	}

	if authReqData.AuthToken == consts.EmptyString {
		return -1, nil
	}

	pInfo, err := srv.components.AuthManager.GetProfileInfo(
		authReqData.AuthToken, consts.EmptyString)
	if err != nil {
		srv.writeWebSockError(conn, err)
		return 0, err
	}

	return pInfo.ID, nil
}
