package objects

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
)

type GameObject interface {
	Type() ObjectType
	Collider() collider.Collider

	Start()
	Update(timeDiff physics.Time)
}

type CollisionChecker func(t transform.Transform, c collider.Collider) GameObject
