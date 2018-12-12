package weakwalls

import (
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/walls"
)

const (
	Type = walls.Type + ".weak"
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
