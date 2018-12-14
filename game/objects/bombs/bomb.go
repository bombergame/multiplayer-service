package bombs

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/objects/bombs/state"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

const (
	DefaultExplosionRadius  = 2
	DefaultExplosionTimeout = 3 * time.Second
)

type Bomb struct {
	objectID   objects.ObjectID
	objectType objects.ObjectType
	state      bombstate.State

	transform components.Transform

	explosionRadius  physics.Integer
	explosionTimeout time.Duration

	changeHandler objects.ChangeHandler
}

func NewBomb() *Bomb {
	return &Bomb{
		explosionRadius:  DefaultExplosionRadius,
		explosionTimeout: DefaultExplosionTimeout,
	}
}

func (b *Bomb) ObjectType() objects.ObjectType {
	return b.objectType
}

func (b *Bomb) SetObjectType(t objects.ObjectType) {
	b.objectType = t
}

func (b *Bomb) ObjectID() objects.ObjectID {
	return b.objectID
}

func (b *Bomb) SetObjectID(id objects.ObjectID) {
	b.objectID = id
}

func (b *Bomb) Transform() components.Transform {
	return b.transform
}

func (b *Bomb) Spawn(pos physics.PositionVec2D) {
	b.state = bombstate.Placed
	b.transform.Position = pos

	b.changeHandler(b)
}

func (b *Bomb) Update(d time.Duration) {
	//TODO
}

func (b *Bomb) SetChangeHandler(h objects.ChangeHandler) {
	b.changeHandler = h
}

//easyjson:json
type MessageData struct {
	objects.MessageData
	State     bombstate.State      `json:"state"`
	Transform components.Transform `json:"transform"`
}

func (b *Bomb) Serialize() interface{} {
	return MessageData{
		MessageData: objects.MessageData{
			ObjectID:   b.objectID,
			ObjectType: b.objectType,
		},
		State:     b.state,
		Transform: b.transform,
	}
}
