package bombs

import (
	"github.com/bombergame/multiplayer-service/game/objects"
)

type Bomb struct {
	objectID   objects.ObjectID
	objectType objects.ObjectType
}
