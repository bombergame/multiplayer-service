package players

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
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

	objectID objects.ID
	state    playerstate.State

	transform transform.Transform

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

func (p *Player) ObjectID() objects.ID {
	return p.objectID
}

func (p *Player) OutChan() *ws.OutChan {
	return p.outChan
}

func (p *Player) SetID(id int64) {
	p.id = id
}

func (p *Player) SetObjectID(id objects.ID) {
	p.objectID = id
}

func (p *Player) Type() objects.ObjectType {
	return Type
}

func (p *Player) Transform() transform.Transform {
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
}

func (p *Player) Update(duration time.Duration) {
	p.handleCommands()
}

func (p *Player) SetChangeHandler(h objects.ChangeHandler) {
	p.changeHandler = h
}

//easyjson:json
type MessageData struct {
	ObjectID  objects.ID          `json:"object_id"`
	State     playerstate.State   `json:"state"`
	Transform transform.Transform `json:"transform"`
}

func (p *Player) GetMessageData() MessageData {
	return MessageData{
		ObjectID:  p.objectID,
		State:     p.state,
		Transform: p.transform,
	}
}

func (p *Player) Serialize() (objects.ObjectType, interface{}) {
	return Type, p.GetMessageData()
}

func (p *Player) SetCellObjectGetter(getter objects.CellObjectGetter) {
	p.objGetter = getter
}

func (p *Player) moveTo(pos physics.PositionVec2D) {
	obj, err := p.objGetter(pos)
	if err != nil || obj != nil {
		return
	}
	p.transform.Position = pos
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
	MovementStep = 1
)

func (p *Player) handleCmd(c playercommands.Cmd) {
	switch c {
	case playercommands.MoveUp:
		p.moveTo(p.transform.Position.Up(MovementStep))

	case playercommands.MoveDown:
		p.moveTo(p.transform.Position.Down(MovementStep))

	case playercommands.MoveLeft:
		p.moveTo(p.transform.Position.Left(MovementStep))

	case playercommands.MoveRight:
		p.moveTo(p.transform.Position.Right(MovementStep))
	}
}
