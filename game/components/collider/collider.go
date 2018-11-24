package collider

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
)

type Collider struct {
	Transform transform.Transform
	Radius    physics.Float
}
