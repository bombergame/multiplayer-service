package wall

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/physics"
)

type Wall struct {
	transform transform.Transform
	collider  collider.Collider
}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) Collider() collider.Collider {
	return w.collider
}

func (w *Wall) Start() {
	//TODO
}

func (w *Wall) Update(timeDiff physics.Time) {
	//TODO
}

func (w *Wall) spawn() {
	//TODO
}
