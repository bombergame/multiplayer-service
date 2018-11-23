package player

import (
	"github.com/bombergame/multiplayer-service/game/physics"
)

const (
	StateAlive = "player.state.alive"
	StateDead  = "player.state.dead"

	DefaultSpeed = 4.0
)

type CommandsChan <-chan Command
type BeforeMoveFunc func(pNew physics.PositionVec2D) error

type Player struct {
	id    int64
	state state

	transform transform

	commands   *CommandsChan
	beforeMove BeforeMoveFunc
}

type state string

type transform struct {
	speed       physics.Speed
	speedVec    physics.SpeedVec2D
	positionVec physics.PositionVec2D
}

func NewPlayer(id int64) *Player {
	return &Player{
		id:    id,
		state: StateAlive,

		transform: transform{
			speed:       DefaultSpeed,
			speedVec:    physics.GetSpeedVec2DZeros(),
			positionVec: physics.GetPositionVec2DZeros(),
		},

		commands:   nil,
		beforeMove: nil,
	}
}

func (p Player) GetID() int64 {
	return p.id
}

func (p *Player) SetCommandsChan(cmdCh *CommandsChan) {
	p.commands = cmdCh
}

func (p *Player) SetBeforeMoveFunc(f BeforeMoveFunc) {
	p.beforeMove = f
}

func (p *Player) MoveTo(pVecNew physics.PositionVec2D) {
	p.tryMoveOrStop(pVecNew)
}

func (p *Player) PerformStep(t physics.Time) {
	if p.state != StateDead {
		p.handleCommands()
		p.handleMovement(t)
	}
}

func (p *Player) handleCommands() {
	for {
		select {
		case c := <-(*p.commands):
			p.handleCommand(c)
		default:
			return
		}
	}
}

func (p *Player) handleCommand(c Command) {
	switch c {
	case CommandStop:
		p.transform.speedVec = physics.GetSpeedVec2DZeros()
	case CommandMoveUp:
		p.transform.speedVec = physics.GetSpeedVec2D(0, DefaultSpeed)
	case CommandMoveDown:
		p.transform.speedVec = physics.GetSpeedVec2D(0, -DefaultSpeed)
	case CommandMoveLeft:
		p.transform.speedVec = physics.GetSpeedVec2D(-DefaultSpeed, 0)
	case CommandMoveRight:
		p.transform.speedVec = physics.GetSpeedVec2D(DefaultSpeed, 0)
	}
}

func (p *Player) handleMovement(t physics.Time) {
	pTempVec := p.transform.positionVec
	pTempVec.Translate(p.transform.speedVec, t)
	p.tryMoveOrStop(pTempVec)
}

func (p *Player) tryMoveOrStop(pVecNew physics.PositionVec2D) {
	if err := p.beforeMove; err == nil {
		p.transform.positionVec = pVecNew
	} else {
		p.transform.speedVec = physics.GetSpeedVec2DZeros()
	}
}
