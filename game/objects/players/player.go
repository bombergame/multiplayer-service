package players

import (
	"github.com/bombergame/multiplayer-service/game/components/collider"
	"github.com/bombergame/multiplayer-service/game/components/movement"
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players/commands"
	"github.com/bombergame/multiplayer-service/game/objects/players/state"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/bombergame/multiplayer-service/utils/ws"
)

type Player struct {
	id    int64
	state state.State

	movement  movement.Movement
	transform transform.Transform
	collider  collider.Collider

	cmdChan *commands.Chan
	outChan *ws.OutChan

	checkCollision objects.CollisionChecker
}

func NewPlayer(id int64) *Player {
	return &Player{
		id: id,
	}
}

func (p *Player) ID() int64 {
	return p.id
}

func (p *Player) OutChan() *ws.OutChan {
	return p.outChan
}

func (p *Player) SetID(id int64) {
	p.id = id
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

func (p *Player) SetOutChan(outChan *ws.OutChan) {
	p.outChan = outChan
}

func (p *Player) SetCollisionChecker(f objects.CollisionChecker) {
	p.checkCollision = f
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
	prevPosVec := p.transform.Position
	p.transform.Position.Translate(p.movement.SpeedVec, timeDiff)

	if p.checkCollision(p.transform, p.collider) != nil {
		p.transform.Position = prevPosVec
		p.movement.SpeedVec = physics.GetSpeedVec2DZeros()
	}
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
		p.movement.SpeedVec = physics.GetSpeedVec2D(0, p.movement.Speed)
	case commands.MoveDown:
		p.movement.SpeedVec = physics.GetSpeedVec2D(0, -p.movement.Speed)
	case commands.MoveLeft:
		p.movement.SpeedVec = physics.GetSpeedVec2D(-p.movement.Speed, 0)
	case commands.MoveRight:
		p.movement.SpeedVec = physics.GetSpeedVec2D(p.movement.Speed, 0)
	}
}
