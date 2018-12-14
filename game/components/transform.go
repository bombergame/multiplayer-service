package components

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

//easyjson:json
type Transform struct {
	Position physics.PositionVec2D `json:"position"`
}
