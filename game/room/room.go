package room

import (
	"github.com/bombergame/multiplayer-service/errs"
	"github.com/bombergame/multiplayer-service/game/physics"
	"github.com/bombergame/multiplayer-service/game/player"
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

	mutex *sync.Mutex
}

func NewRoom(id uuid.UUID, maxPlayers int32) *Room {
	return &Room{
		id:            id,
		ticker:        time.NewTicker(time.Second / TicksPerSecond),
		players:       make(map[int64]*player.Player),
		maxNumPlayers: maxPlayers,
	}
}

func (r *Room) AddPlayer(p *player.Player) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if len(r.players) == int(r.maxNumPlayers) {
		return errs.FullRoomError
	}

	r.players[p.GetID()] = p
	//TODO: Move to random empty cell, add callbacks

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
