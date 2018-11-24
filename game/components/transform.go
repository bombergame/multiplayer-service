package components

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

type Transform struct {
	Position physics.PositionVec2D
}

type CollisionChecker func(transform Transform) physics.Collision
