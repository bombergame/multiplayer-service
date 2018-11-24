package objects

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/mailru/easyjson"
)

type GameObject interface {
	Type() ObjectType
	Collider() collider.Collider

	Start()
	Update(timeDiff physics.Time)

	easyjson.Marshaler
}

type CollisionChecker func(t transform.Transform, c collider.Collider) GameObject
