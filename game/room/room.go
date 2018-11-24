package room

import (
	"github.com/satori/go.uuid"
)

type Room struct {
	id uuid.UUID
}

func NewRoom(id uuid.UUID) *Room {
	return &Room{}
}

func (r *Room) Id() uuid.UUID {
	return r.id
}
