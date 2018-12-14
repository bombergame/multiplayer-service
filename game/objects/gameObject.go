package objects

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components"
	"github.com/bombergame/multiplayer-service/game/errs"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

type ChangeHandler func(GameObject)
type CollisionHandler func(GameObject)

type GameObject interface {
	ObjectType() ObjectType
	SetObjectType(ObjectType)

	ObjectID() ObjectID
	SetObjectID(ObjectID)

	Transform() components.Transform

	Spawn(physics.PositionVec2D)
	Update(time.Duration)

	SetChangeHandler(ChangeHandler)

	Serialize() interface{}
}

const (
	MessageType = "object"
)

//easyjson:json
type MessageData struct {
	ObjectID   ObjectID   `json:"object_id"`
	ObjectType ObjectType `json:"object_type"`
}

type CellObjectGetter func(d physics.PositionVec2D) (GameObject, *errs.InvalidCellIndexError)

type MovingObject interface {
	SetCellObjectGetter(CellObjectGetter)
}

type WeakObject interface {
	Destroy()
}
