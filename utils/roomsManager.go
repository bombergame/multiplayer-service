package utils

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/rooms"
	"github.com/satori/go.uuid"
	"sync"
)

type RoomsManager struct {
	rwMutex *sync.RWMutex
	rooms   map[uuid.UUID]*rooms.Room
}

func NewRoomsManager() *RoomsManager {
	return &RoomsManager{
		rwMutex: &sync.RWMutex{},
		rooms:   make(map[uuid.UUID]*rooms.Room, 0),
	}
}

func (rm *RoomsManager) AddRoom(r *rooms.Room) error {
	rm.rwMutex.Lock()
	defer rm.rwMutex.Unlock()

	if _, ok := rm.rooms[r.ID()]; ok {
		return errs.NewDuplicateError("room already exists")
	}

	rm.rooms[r.ID()] = r
	return nil
}

func (rm *RoomsManager) GetRoom(id uuid.UUID) (*rooms.Room, error) {
	rm.rwMutex.RLock()
	defer rm.rwMutex.RUnlock()

	r, ok := rm.rooms[id]
	if !ok {
		return nil, errs.NewNotFoundError("room not found")
	}

	return r, nil
}
