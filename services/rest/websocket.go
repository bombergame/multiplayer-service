package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"sync"
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

	var msg ws.InMessage
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

	srv.Logger().Info("Auth message received")

	authID, err := srv.handleAuthRequest(conn, &msg)
	if err != nil {
		return
	}

	srv.Logger().Info("Auth message handled")

	room, err := srv.components.RoomsManager.GetRoom(roomID)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	srv.Logger().Info("Auth message received")

	outChan := make(ws.OutChan, 10)

	p := players.NewPlayer(authID)
	p.SetOutChan(&outChan)

	if err := room.AddPlayer(p); err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			select {
			case c := <-outChan:
				srv.writeWebSockJSON(conn, c)
			}
		}
	}()

	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				srv.closeConnectionWithError(conn, err)
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}

func (srv *Service) handleAuthRequest(conn *websocket.Conn, msg *ws.InMessage) (int64, error) {
	var authReqData ws.AuthMessageData

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
