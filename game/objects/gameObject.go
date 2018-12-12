package objects

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

type GameObject interface {
	Type() ObjectType
	ObjectID() ID

	Transform() transform.Transform

	Spawn(physics.PositionVec2D)
	Update(time.Duration)
}
