package weakwalls

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/walls"
	"github.com/bombergame/multiplayer-service/game/objects/walls/weak/state"
)

const (
	Type = walls.Type + ".weak"
)

type Wall struct {
	walls.Wall
	state weakwallstate.State
}

func NewWall() *Wall {
	return &Wall{
		Wall:  *walls.NewWall(),
		state: weakwallstate.Up,
	}
}

func (w *Wall) Type() objects.ObjectType {
	return Type
}

//easyjson:json
type MessageData struct {
	walls.MessageData
	State weakwallstate.State
}

func (w *Wall) GetMessageData() MessageData {
	return MessageData{
		MessageData: w.Wall.GetMessageData(),
		State:       w.state,
	}
}

func (w *Wall) Serialize() interface{} {
	return w.GetMessageData()
}
