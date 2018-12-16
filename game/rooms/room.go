package rooms

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/domains"
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/fields"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/players"
	"github.com/bombergame/multiplayer-service/game/physics"
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
)

const (
	TickerPeriod          = 20 * time.Millisecond
	BroadcastTickerPeriod = time.Second
)

type Room struct {
	id    uuid.UUID
	title string

	state gamestate.State

	tLimit time.Duration
	ticker *time.Ticker

	field   *fields.Field
	objects map[objects.ObjectID]objects.GameObject

	maxNumPlayers   int64
	allowAnonymous  bool
	numAlivePlayers int64
	players         map[int64]*players.Player

	cmdChan gamecommands.CmdChan

	mu sync.RWMutex
}

func NewRoom(r domains.Room) *Room {
	return &Room{
		id:    uuid.Nil,
		title: r.Title,

		state: gamestate.Pending,

		tLimit: time.Duration(r.TimeLimit) * time.Minute,
		ticker: time.NewTicker(TickerPeriod),

		field: fields.NewField(physics.GetSize2D(
			physics.Integer(r.FieldSize.Width),
			physics.Integer(r.FieldSize.Height),
		)),

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

func (r *Room) SetID(id uuid.UUID) {
	r.id = id
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
					r.withLock(r.startGame)
				case gamecommands.Stop:
					r.withLock(r.stopGame)
				case gamecommands.End:
					r.withLock(r.endGame)
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
	p.SetDeathHandler(func(p *players.Player) {
		r.numAlivePlayers--
		if r.numAlivePlayers == 0 {
			r.endGame()
		}
	})

	r.broadcastState()

	return nil
}

func (r *Room) DeletePlayer(p *players.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.players, p.ID())
	r.broadcastState()
}

func (r *Room) withLock(f func()) {
	r.mu.Lock()
	defer r.mu.Unlock()
	f()
}

func (r *Room) startGame() {
	switch r.state {
	case gamestate.Pending:
		r.state = gamestate.On
		r.broadcastState()

		r.numAlivePlayers = int64(len(r.players))

		r.field.PlaceObjects(r.players)
		r.field.SpawnObjects(func(obj objects.GameObject) {
			r.broadcast(ws.OutMessage{
				Type: objects.MessageType,
				Data: obj.Serialize(),
			})
		})

		go r.gameLoop()

	case gamestate.Paused:
		r.state = gamestate.On
		r.broadcastState()

	default:
		return
	}
}

func (r *Room) stopGame() {
	if r.state != gamestate.On {
		return
	}

	r.state = gamestate.Paused
	r.broadcastState()
}

func (r *Room) endGame() {
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

	tPrev := tStart
	tPrevBcst := tStart

	for tCur := range r.ticker.C {
		r.mu.Lock()

		if r.state == gamestate.On {
			t := tCur.Sub(tStart)
			if t > r.tLimit {
				break
			}

			d := tCur.Sub(tPrev)
			tPrev = tCur
			r.field.UpdateObjects(d)

			if tCur.Sub(tPrevBcst) > BroadcastTickerPeriod {
				r.broadcastTicker(t)
				tPrevBcst = tCur
			}
		}

		r.endGame()
		r.mu.Unlock()
	}
}

//easyjson:json
type RoomMessageData struct {
	Title     string         `json:"title"`
	State     string         `json:"state"`
	Players   []int64        `json:"players"`
	FieldSize physics.Size2D `json:"field_size"`
}

//easyjson:json
type TickerMessageData struct {
	Value float64 `json:"value"`
}

func (r *Room) broadcastState() {
	p := make([]int64, 0)
	for id := range r.players {
		p = append(p, id)
	}

	message := ws.OutMessage{
		Type: ws.RoomMessageType,
		Data: RoomMessageData{
			Title:     r.title,
			State:     r.state.ToString(),
			Players:   p,
			FieldSize: r.field.Size(),
		},
	}

	r.broadcast(message)
}

func (r *Room) broadcastTicker(t time.Duration) {
	message := ws.OutMessage{
		Type: ws.TickerMessageType,
		Data: TickerMessageData{
			Value: r.tLimit.Seconds() - t.Seconds(),
		},
	}
	r.broadcast(message)
}

func (r *Room) broadcast(message ws.OutMessage) {
	log.Println(message)
	for _, p := range r.players {
		*p.OutChan() <- message
	}
}
