package player

import (
	"github.com/bombergame/multiplayer-service/game/components"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/player/commands"
)

type Player struct {
	id    int64
	state State

	movement  components.Movement
	transform components.Transform

	commands commands.Chan

	checkCollision components.CollisionChecker
}

func (p *Player) GetType() objects.ObjectType {
	return objects.Player
}

func (p *Player) Update() {

}
