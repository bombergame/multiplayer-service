package rooms

import (
	"github.com/bombergame/multiplayer-service/domains"
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

const (
	DefaultAnonymousPlayerID = -1

	TickerPeriod = 20 * time.Millisecond
)

type BroadcastChan chan ws.OutMessage

type Room struct {
	id    uuid.UUID
	title string

	state  GameState
	ticker *time.Ticker

	maxNumPlayers  int64
	allowAnonymous bool
	players        map[int64]*players.Player

	broadcastChan chan ws.OutMessage

	mu sync.RWMutex
}

func NewRoom(r domains.Room) *Room {
	return &Room{
		id:    r.ID,
		title: r.Title,

		state:  GameStatePending,
		ticker: time.NewTicker(TickerPeriod),

		maxNumPlayers:  r.MaxNumPlayers,
		allowAnonymous: r.AllowAnonymous,
		players:        make(map[int64]*players.Player, 0),

		mu: sync.RWMutex{},
	}
}

func (r *Room) ID() uuid.UUID {
	return r.id
}

func (r *Room) AddPlayer(p *players.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.state != GameStatePending {
		return errs.NewGameError("game already started")
	}

	if int64(len(r.players)) == r.maxNumPlayers {
		return errs.NewGameError("players limit exceeded")
	}

	if p.ID() == DefaultAnonymousPlayerID {
		if !r.allowAnonymous {
			return errs.NewGameError("anonymous not allowed")
		}
		p.SetID(r.findFreeAnonID())
	}

	r.players[p.ID()] = p
	r.broadcastState()

	return nil
}

func (r *Room) DeletePlayer(p *players.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.players, p.ID())
}

func (r *Room) StartGame() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.state != GameStatePending {
		return
	}

	r.state = GameStateOn
	r.broadcastState()

	go r.gameLoop()
}

func (r *Room) StopGame() {
	//TODO
}

func (r *Room) findFreeAnonID() int64 {
	var i int64
	for i = 1; i < r.maxNumPlayers; i++ {
		if _, ok := r.players[i]; !ok {
			return i
		}
	}
	return 0
}

func (r *Room) broadcastState() {
	p := make([]int64, 0)
	for id := range r.players {
		p = append(p, id)
	}

	message := ws.OutMessage{
		Type: "game",
		Data: ws.RoomMessageData{
			Title:   r.title,
			State:   r.state.ToString(),
			Players: p,
		},
	}

	r.broadcast(message)
}

func (r *Room) broadcast(message ws.OutMessage) {
	log.Println("Broadcast message: ", message)
	for _, p := range r.players {
		*p.OutChan() <- message
	}
}

func (r *Room) gameLoop() {
	for {
		break //TODO
	}
}
