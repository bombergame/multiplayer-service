package movement

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

type Movement struct {
	Speed    physics.Speed
	SpeedVec physics.SpeedVec2D
}
