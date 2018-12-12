package rooms

import (
	"github.com/bombergame/multiplayer-service/domains"
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/fields"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/rooms/commands"
	"github.com/bombergame/multiplayer-service/game/rooms/state"
	"github.com/bombergame/multiplayer-service/utils/ws"
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

const (
	DefaultAnonymousPlayerID = -1

	TickerPeriod = 50 * time.Millisecond
)

type Room struct {
	id    uuid.UUID
	title string

	state gamestate.State

	tLimit time.Duration
	ticker *time.Ticker

	field   *fields.Field
	objects map[objects.ID]objects.GameObject

	maxNumPlayers  int64
	allowAnonymous bool
	players        map[int64]*players.Player

	cmdChan gamecommands.CmdChan

	mu sync.RWMutex
}

func NewRoom(r domains.Room) *Room {
	return &Room{
		id:    r.ID,
		title: r.Title,

		state: gamestate.Pending,

		tLimit: time.Duration(r.TimeLimit) * time.Minute,
		ticker: time.NewTicker(TickerPeriod),

		field: fields.NewField(fields.Size{
			Width:  r.FieldSize.Width,
			Height: r.FieldSize.Height,
		}),

		maxNumPlayers:  r.MaxNumPlayers,
		allowAnonymous: r.AllowAnonymous,
		players:        make(map[int64]*players.Player, 0),

		cmdChan: make(gamecommands.CmdChan, gamecommands.ChanLen),

		mu: sync.RWMutex{},
	}
}

func (r *Room) ID() uuid.UUID {
	return r.id
}

func (r *Room) CmdChan() *gamecommands.CmdChan {
	return &r.cmdChan
}

func (r *Room) RunGame() {
	go func() {
		for {
			select {
			case c := <-r.cmdChan:
				switch c {
				case gamecommands.Start:
					r.startGame()
				case gamecommands.Stop:
					r.stopGame()
				case gamecommands.End:
					r.endGame()
				}
			}
		}
	}()
}

func (r *Room) AddPlayer(p *players.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.state != gamestate.Pending {
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

func (r *Room) startGame() {
	r.mu.Lock()
	defer r.mu.Unlock()

	switch r.state {
	case gamestate.Pending:
		r.state = gamestate.On

		r.field.SpawnObjects(int32(r.maxNumPlayers))
		go r.gameLoop()

	case gamestate.Paused:
		r.state = gamestate.On

	default:
		return
	}

	r.broadcastState()
}

func (r *Room) stopGame() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.state != gamestate.On {
		return
	}

	r.state = gamestate.Paused
	r.broadcastState()
}

func (r *Room) endGame() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.state = gamestate.Off
	r.broadcastState()
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

func (r *Room) gameLoop() {
	tStart := <-r.ticker.C
	//tPrev := tStart //TODO

	for tCur := range r.ticker.C {
		r.mu.Lock()

		if r.state == gamestate.On {
			t := tCur.Sub(tStart)
			if t > r.tLimit {
				break
			}

			r.broadcastTicker(t)
		}

		r.mu.Unlock()
	}
}

func (r *Room) broadcastState() {
	p := make([]int64, 0)
	for id := range r.players {
		p = append(p, id)
	}

	message := ws.OutMessage{
		Type: ws.RoomMessageType,
		Data: ws.RoomMessageData{
			Title:   r.title,
			State:   r.state.ToString(),
			Players: p,
		},
	}

	r.broadcast(message)
}

func (r *Room) broadcastTicker(t time.Duration) {
	message := ws.OutMessage{
		Type: ws.TickerMessageType,
		Data: ws.TickerMessageData{
			Value: r.tLimit.Seconds() - t.Seconds(),
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
