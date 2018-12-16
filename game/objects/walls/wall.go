package walls

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/components"
	"github.com/bombergame/multiplayer-service/game/objects"
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

const (
	Type = "wall"
)

type Wall struct {
	objectID   objects.ObjectID
	objectType objects.ObjectType

	transform     components.Transform
	changeHandler objects.ChangeHandler
}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) ObjectID() objects.ObjectID {
	return w.objectID
}

func (w *Wall) ObjectType() objects.ObjectType {
	return w.objectType
}

func (w *Wall) SetObjectID(id objects.ObjectID) {
	w.objectID = id
}

func (w *Wall) SetObjectType(t objects.ObjectType) {
	w.objectType = t
}

func (w *Wall) Transform() components.Transform {
	return w.transform
}

func (w *Wall) Spawn(position physics.PositionVec2D) {
	w.transform.Position = position
	w.changeHandler(w)
}

func (w *Wall) Update(duration time.Duration) {
}

func (w *Wall) ChangeHandler() objects.ChangeHandler {
	return w.changeHandler
}

func (w *Wall) SetChangeHandler(h objects.ChangeHandler) {
	w.changeHandler = h
}

//easyjson:json
type MessageData struct {
	objects.MessageData
	Transform components.Transform `json:"transform"`
}

func (w *Wall) GetMessageData() MessageData {
	return MessageData{
		MessageData: objects.MessageData{
			ObjectID:   w.objectID,
			ObjectType: w.objectType,
		},
		Transform: w.transform,
	}
}

func (w *Wall) Serialize() interface{} {
	return w.GetMessageData()
}
