package walls

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/objects"
)

type Wall struct {
	objType   objects.ObjectType
	transform transform.Transform
}
