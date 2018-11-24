package wall

import (
	"github.com/bombergame/multiplayer-service/game/objects"
)

type WeakWall struct {
	Wall
}

func (w *WeakWall) Type() objects.ObjectType {
	return objects.WallWeak
}
