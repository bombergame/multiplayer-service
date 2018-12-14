package solidwalls

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/walls"
)

const (
	Type = walls.Type + ".solid"
)

type Wall struct {
	walls.Wall
}

func NewWall() *Wall {
	return &Wall{
		Wall: *walls.NewWall(),
	}
}

func (w *Wall) Type() objects.ObjectType {
	return Type
}

//easyjson:json
type MessageData struct {
	walls.MessageData
}

func (w *Wall) GetMessageData() MessageData {
	return MessageData{
		MessageData: w.Wall.GetMessageData(),
	}
}

func (w *Wall) Serialize() (objects.ObjectType, interface{}) {
	return Type, w.GetMessageData()
}
