package roommessages

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/rooms/state"
	"github.com/satori/go.uuid"
)

const (
	MessageType    = "room"
	NewPlayerEvent = "newplayer"
)

//easyjson:json
type MessageData struct {
	ID        uuid.UUID       `json:"id"`
	GameState gamestate.State `json:"game_state"`
}

//easyjson:json
type NewPlayerMessageData struct {
}
