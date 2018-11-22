package objects

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

type GameObject interface {
	PerformStep(t physics.Time)
}
