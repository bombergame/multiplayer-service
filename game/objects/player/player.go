package player

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/movement"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/player/commands"
	"github.com/bombergame/multiplayer-service/game/objects/player/state"
	"github.com/bombergame/multiplayer-service/game/physics"
)

const (
	DefaultSpeed = 0.05
)

type Player struct {
	id    int64
	state state.State

	movement  movement.Movement
	transform transform.Transform
	collider  collider.Collider

	cmdChan *commands.Chan

	checkCollision objects.CollisionChecker
}

func (p *Player) Type() objects.ObjectType {
	return objects.Player
}

func (p *Player) Transform() transform.Transform {
	return p.transform
}

func (p *Player) SetCmdChan(cmdChan *commands.Chan) {
	p.cmdChan = cmdChan
}

func (p *Player) Start() {
	p.spawn()
	p.state = state.Alive
}

func (p *Player) Update(timeDiff physics.Time) {
	p.handleCommands()
	p.move(timeDiff)
}

func (p *Player) spawn() {
	//TODO
}

func (p *Player) move(timeDiff physics.Time) {
	//TODO
}

func (p *Player) handleCommands() {
	for {
		select {
		case c := <-(*p.cmdChan):
			p.handleCommand(c)
		default:
			return
		}
	}
}

func (p *Player) handleCommand(c commands.Command) {
	switch c {
	case commands.Stop:
		p.movement.SpeedVec = physics.GetSpeedVec2DZeros()
	case commands.MoveUp:
		p.movement.SpeedVec = physics.GetSpeedVec2D(0, DefaultSpeed)
	case commands.MoveDown:
		p.movement.SpeedVec = physics.GetSpeedVec2D(0, -DefaultSpeed)
	case commands.MoveLeft:
		p.movement.SpeedVec = physics.GetSpeedVec2D(-DefaultSpeed, 0)
	case commands.MoveRight:
		p.movement.SpeedVec = physics.GetSpeedVec2D(DefaultSpeed, 0)
	}
}
