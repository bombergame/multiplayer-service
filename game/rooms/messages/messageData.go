package roommessages

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/rooms/events"
	"github.com/satori/go.uuid"
)

const (
	MessageType = "room"
)

//easyjson:json
type MessageData struct {
	ID uuid.UUID `json:"id"`
}

//easyjson:json
type InfoMessageData struct {
	MessageData
}

//easyjson:json
type EventMessageData struct {
	MessageData
	Event roomevents.Event `json:"event"`
}

//easyjson:json
type NewPlayerMessageData struct {
	ProfileID int64 `json:"id"`
}
