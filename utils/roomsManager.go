package utils

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/room"
	"github.com/satori/go.uuid"
	"sync"
)

type RoomsManager struct {
	rwMutex *sync.RWMutex
	rooms   map[uuid.UUID]*room.Room
}

func NewRoomsManager() *RoomsManager {
	return &RoomsManager{
		rwMutex: &sync.RWMutex{},
		rooms:   make(map[uuid.UUID]*room.Room, 0),
	}
}

func (rm *RoomsManager) AddRoom(r *room.Room) error {
	rm.rwMutex.Lock()
	defer rm.rwMutex.Unlock()

	_, ok := rm.rooms[r.Id()]
	if ok {
		return errs.NewDuplicateError("room already exists")
	}

	rm.rooms[r.Id()] = r
	return nil
}

func (rm *RoomsManager) GetRoom(id uuid.UUID) (*room.Room, error) {
	rm.rwMutex.RLock()
	defer rm.rwMutex.RUnlock()

	r, ok := rm.rooms[id]
	if ok {
		return nil, errs.NewNotFoundError("room not found")
	}

	return r, nil
}

func (rm *RoomsManager) DeleteRoom(id uuid.UUID) error {
	rm.rwMutex.Lock()
	defer rm.rwMutex.Unlock()

	_, ok := rm.rooms[id]
	if ok {
		return errs.NewNotFoundError("room not found")
	}

	delete(rm.rooms, id)
	return nil
}
