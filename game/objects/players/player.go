package players

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players/commands"
	"github.com/bombergame/multiplayer-service/game/objects/players/state"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"sync"
	"time"
)

const (
	Type = "player"
)

type Player struct {
	id int64

	objectID   objects.ObjectID
	objectType objects.ObjectType
	state      playerstate.State

	movement  components.Movement
	transform components.Transform

	cmdChan *playercommands.CmdChan
	outChan *ws.OutChan

	objGetter     objects.CellObjectGetter
	changeHandler objects.ChangeHandler

	mu *sync.Mutex
}

func NewPlayer(id int64) *Player {
	return &Player{
		id: id,
		mu: &sync.Mutex{},
	}
}

func (p *Player) ID() int64 {
	return p.id
}

func (p *Player) ObjectID() objects.ObjectID {
	return p.objectID
}

func (p *Player) ObjectType() objects.ObjectType {
	return p.objectType
}

func (p *Player) OutChan() *ws.OutChan {
	return p.outChan
}

func (p *Player) SetID(id int64) {
	p.id = id
}

func (p *Player) SetObjectID(id objects.ObjectID) {
	p.objectID = id
}

func (p *Player) SetObjectType(t objects.ObjectType) {
	p.objectType = t
}

func (p *Player) Transform() components.Transform {
	return p.transform
}

func (p *Player) SetCmdChan(cmdChan *playercommands.CmdChan) {
	p.cmdChan = cmdChan
}

func (p *Player) SetOutChan(outChan *ws.OutChan) {
	p.outChan = outChan
}

func (p *Player) Spawn(pos physics.PositionVec2D) {
	p.state = playerstate.Alive

	p.transform.Position = pos

	p.movement.StepSize = DefaultStepSize
	p.movement.MinStepInterval = DefaultMinStepInterval
	p.movement.LastStepInterval = 2 * p.movement.MinStepInterval

	p.changeHandler(p)
}

func (p *Player) Update(duration time.Duration) {
	p.passTime(duration)
	p.handleCommands()
}

func (p *Player) SetChangeHandler(h objects.ChangeHandler) {
	p.changeHandler = h
}

//easyjson:json
type MessageData struct {
	objects.MessageData
	ID        int64                          `json:"id"`
	Transform components.Transform           `json:"transform"`
	Movement  components.MovementMessageData `json:"movement"`
}

func (p *Player) GetMessageData() MessageData {
	return MessageData{
		MessageData: objects.MessageData{
			ObjectID:   p.objectID,
			ObjectType: p.objectType,
		},
		ID:        p.id,
		Transform: p.transform,
		Movement:  p.movement.GetMessageData(),
	}
}

func (p *Player) Serialize() interface{} {
	return p.GetMessageData()
}

func (p *Player) SetCellObjectGetter(getter objects.CellObjectGetter) {
	p.objGetter = getter
}

func (p *Player) move(newPos physics.PositionVec2D) {
	if p.movement.LastStepInterval < p.movement.MinStepInterval {
		return
	}
	obj, err := p.objGetter(newPos)
	if err != nil || obj != nil {
		return
	}
	p.transform.Position = newPos
	p.movement.LastStepInterval = 0
	p.changeHandler(p)
}

func (p *Player) handleCommands() {
	for {
		select {
		case c := <-(*p.cmdChan):
			p.handleCmd(c)
		default:
			return
		}
	}
}

const (
	DefaultStepSize        = 1
	DefaultMinStepInterval = time.Second / 2
)

func (p *Player) handleCmd(c playercommands.Cmd) {
	switch c {
	case playercommands.MoveUp:
		p.move(p.transform.Position.Up(p.movement.StepSize))

	case playercommands.MoveDown:
		p.move(p.transform.Position.Down(p.movement.StepSize))

	case playercommands.MoveLeft:
		p.move(p.transform.Position.Left(p.movement.StepSize))

	case playercommands.MoveRight:
		p.move(p.transform.Position.Right(p.movement.StepSize))
	}
}

func (p *Player) passTime(duration time.Duration) {
	p.movement.LastStepInterval += duration
}
