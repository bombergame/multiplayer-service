package rooms

import (
	"github.com/satori/go.uuid"
)

type Room struct {
	id uuid.UUID
}

func NewRoom(id uuid.UUID) *Room {
	return &Room{
		id: id,
	}
}

func (r *Room) Id() uuid.UUID {
	return r.id
}
