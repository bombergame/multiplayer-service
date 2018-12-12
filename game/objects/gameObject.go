package objects

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

type ChangeHandler func(obj GameObject)

type GameObject interface {
	Type() ObjectType

	ObjectID() ID
	SetObjectID(ID)

	Transform() transform.Transform

	Spawn(physics.PositionVec2D)
	Update(time.Duration)

	SetChangeHandler(ChangeHandler)

	Serialize() interface{}
}
