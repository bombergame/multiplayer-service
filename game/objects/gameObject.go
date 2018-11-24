package objects

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

type GameObject interface {
	GetType() ObjectType
	Update(timeDiff physics.Time)
}
