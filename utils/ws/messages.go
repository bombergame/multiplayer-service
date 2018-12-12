package ws

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/objects"
)

//easyjson:json
type InMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

//easyjson:json
type OutMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type OutChan chan OutMessage

const (
	OutChanLen = 10

	AuthMessageType = "auth"

	RoomMessageType   = "room"
	FieldMessageType  = "field"
	TickerMessageType = "ticker"

	OkMessageType    = "ok"
	ErrorMessageType = "error"
)

//easyjson:json
type AuthMessageData struct {
	AuthToken string `json:"auth_token"`
	UserAgent string `json:"user_agent"`
}

//easyjson:json
type RoomMessageData struct {
	Title   string  `json:"title"`
	State   string  `json:"state"`
	Players []int64 `json:"players"`
}

//easyjson:json
type TickerMessageData struct {
	Value float64 `json:"value"`
}

//easyjson:json
type WallMessageData struct {
	Cells [][]objects.ObjectType `json:"cells"`
}

//easyjson:json
type OkMessageData struct {
	Message string `json:"message"`
}

//easyjson:json
type ErrorMessageData struct {
	Message string `json:"message"`
}
