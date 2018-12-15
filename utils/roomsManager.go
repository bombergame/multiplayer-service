package utils

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/multiplayer-service/game/rooms"
	"github.com/satori/go.uuid"
	"sync"
)

const (
	RoomNotFoundMessage = "room not found"
)

type RoomsManager struct {
	rwMutex     *sync.RWMutex
	rooms       map[uuid.UUID]*rooms.Room
	numRooms    int64
	notFoundErr *errs.NotFoundError
}

func NewRoomsManager() *RoomsManager {
	return &RoomsManager{
		rwMutex:     &sync.RWMutex{},
		rooms:       make(map[uuid.UUID]*rooms.Room, 0),
		numRooms:    0,
		notFoundErr: errs.NewNotFoundError(RoomNotFoundMessage).(*errs.NotFoundError),
	}
}

func (rm *RoomsManager) AddRoom(r *rooms.Room) (int64, error) {
	rm.rwMutex.Lock()
	defer rm.rwMutex.Unlock()

	for {
		id := uuid.NewV4()
		if _, ok := rm.rooms[id]; !ok {
			r.SetID(id)
			break
		}
	}

	rm.rooms[r.ID()] = r
	rm.numRooms++

	return rm.numRooms, nil
}

func (rm *RoomsManager) GetRoom(id uuid.UUID) (*rooms.Room, error) {
	rm.rwMutex.RLock()
	defer rm.rwMutex.RUnlock()

	r, ok := rm.rooms[id]
	if !ok {
		return nil, rm.notFoundErr
	}

	return r, nil
}

func (rm *RoomsManager) DeleteRoom(id uuid.UUID) (int64, error) {
	rm.rwMutex.Lock()
	defer rm.rwMutex.Unlock()

	_, ok := rm.rooms[id]
	if !ok {
		return rm.numRooms, rm.notFoundErr
	}

	delete(rm.rooms, id)
	rm.numRooms--

	return rm.numRooms, nil
}
