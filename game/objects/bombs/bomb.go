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
	Type = "bomb"
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

	changeHandler    objects.ChangeHandler
	explosionHandler objects.ExplosionHandler
}

func NewBomb() *Bomb {
	return &Bomb{
		state: bombstate.Inactive,
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

func (b *Bomb) State() bombstate.State {
	return b.state
}

func (b *Bomb) Transform() components.Transform {
	return b.transform
}

func (b *Bomb) ExplosionRadius() physics.Integer {
	return b.explosionRadius
}

func (b *Bomb) SetExplosionHandler(h objects.ExplosionHandler) {
	b.explosionHandler = h
}

func (b *Bomb) Spawn(pos physics.PositionVec2D) {
	b.state = bombstate.Placed
	b.transform.Position = pos

	b.explosionRadius = DefaultExplosionRadius
	b.explosionTimeout = DefaultExplosionTimeout

	b.changeHandler(b)
}

func (b *Bomb) Update(d time.Duration) {
	if b.state != bombstate.Placed {
		return
	}
	b.explosionTimeout -= d
	if b.explosionTimeout < 0 {
		b.detonate()
	}
}

func (b *Bomb) SetChangeHandler(h objects.ChangeHandler) {
	b.changeHandler = h
}

func (b *Bomb) Collapse() {
	b.detonate()
}

//easyjson:json
type MessageData struct {
	objects.MessageData
	State            bombstate.State      `json:"state"`
	Transform        components.Transform `json:"transform"`
	ExplosionRadius  physics.Integer      `json:"explosion_radius"`
	ExplosionTimeout float64              `json:"explosion_timeout"`
}

func (b *Bomb) Serialize() interface{} {
	return MessageData{
		MessageData: objects.MessageData{
			ObjectID:   b.objectID,
			ObjectType: b.objectType,
		},
		State:            b.state,
		Transform:        b.transform,
		ExplosionRadius:  b.explosionRadius,
		ExplosionTimeout: b.explosionTimeout.Seconds(),
	}
}

func (b *Bomb) detonate() {
	b.state = bombstate.Detonated
	b.changeHandler(b)
	b.explosionHandler(b)
}
