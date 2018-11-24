package wall

import (
	"github.com/bombergame/multiplayer-service/game/objects"
)

type SolidWall struct {
	Wall
}

func (w *SolidWall) Type() objects.ObjectType {
	return objects.WallSolid
}
