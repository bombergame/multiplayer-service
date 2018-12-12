package walls

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components/transform"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

const (
	Type = "wall"
)

type Wall struct {
	objectID      objects.ID
	transform     transform.Transform
	changeHandler objects.ChangeHandler
}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) Type() objects.ObjectType {
	return Type
}

func (w *Wall) ObjectID() objects.ID {
	return w.objectID
}

func (w *Wall) SetObjectID(id objects.ID) {
	w.objectID = id
}

func (w *Wall) Transform() transform.Transform {
	return w.transform
}

func (w *Wall) Spawn(position physics.PositionVec2D) {
	w.transform.Position = position
	w.changeHandler(w)
}

func (w *Wall) Update(duration time.Duration) {
	//TODO
}

func (w *Wall) SetChangeHandler(h objects.ChangeHandler) {
	w.changeHandler = h
}

//easyjson:json
type MessageData struct {
	ObjectID  objects.ID          `json:"object_id"`
	Transform transform.Transform `json:"transform"`
}

func (w *Wall) Serialize() interface{} {
	return MessageData{
		ObjectID:  w.objectID,
		Transform: w.transform,
	}
}
