package rest

import (
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

	var msg WebSocketMessage
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

	_, err = srv.components.RoomsManager.GetRoom(roomID)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	for {
		_ = err
	}
}

//func (srv *Service) handleAuthRequest(conn *websocket.Conn, msg *WebSocketMessage) (int64, error) {
//	var authReqData AuthRequestData
//	if err := easyjson.Unmarshal(msg.Data, &authReqData); err != nil {
//		srv.writeWebSockError(conn, err)
//		return 0, err
//	}
//
//	srv.Logger().Info("auth request data: ", authReqData)
//
//	if authReqData.AuthToken == consts.EmptyString {
//		return -1, nil
//	}
//
//	pInfo, err := srv.components.AuthManager.GetProfileInfo(
//		authReqData.AuthToken, consts.EmptyString)
//	if err != nil {
//		srv.writeWebSockError(conn, err)
//		return 0, err
//	}
//
//	return pInfo.ID, nil
//	return 0, nil
//}

func (srv *Service) closeConnectionWithError(conn *websocket.Conn, err error) {
	srv.writeWebSockError(conn, err)
	if err := conn.Close(); err != nil {
		srv.Logger().Error(err)
	}
}
