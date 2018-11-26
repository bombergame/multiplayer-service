package rooms

import (
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type Room struct {
	id    uuid.UUID
	title string

	state  GameState
	ticker time.Ticker

	players map[int64]*players.Player

	mu sync.RWMutex
}

func NewRoom(id uuid.UUID) *Room {
	return &Room{
		id: id,
	}
}

func (r *Room) Id() uuid.UUID {
	return r.id
}

func (r *Room) AddPlayer(p *players.Player) error {
	return nil //TODO
}
