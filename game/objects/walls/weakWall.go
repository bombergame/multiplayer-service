package walls

import (
	"github.com/bombergame/multiplayer-service/game/objects"
)

type WeakWall struct {
	Wall
}

func NewWeakWall() *WeakWall {
	return &WeakWall{
		Wall: *NewWall(),
	}
}

func (w *WeakWall) Type() objects.ObjectType {
	return objects.WallWeak
}
