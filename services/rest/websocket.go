package rest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/objects/players/commands"
	"github.com/bombergame/multiplayer-service/game/rooms/commands"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strings"
	"sync"
)

func (srv *Service) handleGameplay(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	roomID, err := srv.readRoomID(r)
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	_, b, err := conn.ReadMessage()
	if err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	var msg ws.InMessage
	if err := easyjson.Unmarshal(b, &msg); err != nil {
		srv.closeConnectionWithError(conn, err)
		return
	}

	if msg.Type != ws.AuthMessageType {
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

	cmdChan := make(playercommands.CmdChan, playercommands.ChanLen)
	p.SetCmdChan(&cmdChan)

	outChan := make(ws.OutChan, ws.OutChanLen)
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
			_, p, err := conn.ReadMessage()
			if err != nil {
				srv.closeConnectionWithError(conn, err)
				wg.Done()
				return
			}

			cmd := string(p)
			if strings.HasPrefix(cmd, gamecommands.Prefix) {
				*room.CmdChan() <- gamecommands.Cmd(cmd)
			}
			if strings.HasPrefix(cmd, playercommands.Prefix) {
				cmdChan <- playercommands.Cmd(cmd)
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
