package room

import (
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects/field"
	"github.com/bombergame/multiplayer-service/game/objects/player"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

const (
	MaxNumPlayers = 4

	GameStateOn     = "game.gameState.on"
	GameStatePaused = "game.gameState.paused"
	GameStateOver   = "game.gameState.over"

	TicksPerSecond = 20
	TicksTimeDiff  = physics.Time(1.0 / TicksPerSecond)
)

type Room struct {
	id        uuid.UUID
	gameState gameState

	ticker *time.Ticker

	maxNumPlayers int32
	players       map[int64]*player.Player

	field *field.Field

	mutex *sync.Mutex
}

func NewRoom(id uuid.UUID, maxPlayers int32) *Room {
	return &Room{
		id:            id,
		ticker:        time.NewTicker(time.Second / TicksPerSecond),
		maxNumPlayers: maxPlayers,
		players:       make(map[int64]*player.Player),
		field:         field.NewField(field.GetSize(field.DefaultWidth, field.DefaultHeight)),
	}
}

func (r *Room) AddPlayer(p *player.Player) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if len(r.players) == int(r.maxNumPlayers) {
		return errs.FullRoomError
	}

	p.SetBeforeMoveFunc(func(pNew physics.PositionVec2D) error {
		if !r.field.IsValidPosition(pNew) || !r.field.IsCellEmpty(pNew) {
			return errs.CannotMoveError
		}
		return nil
	})
	p.MoveTo(r.field.GetRandomEmptyPosition())

	r.players[p.GetID()] = p

	return nil
}

func (r *Room) DeletePlayer(p *player.Player) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.players[p.GetID()]; !ok {
		return errs.PlayerNotFoundError
	}

	delete(r.players, p.GetID())

	return nil
}

func (r *Room) Run() error {
	go r.gameLoop()
	return nil
}

func (r *Room) gameLoop() {
	for range r.ticker.C {
		switch r.gameState {
		case GameStateOn:
			for _, p := range r.players {
				p.PerformStep(TicksTimeDiff)
			}
		}
	}
}

type gameState string
