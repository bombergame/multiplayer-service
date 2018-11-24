package rest

import (
	"github.com/satori/go.uuid"
)

//easyjson:json
type Room struct {
	ID            uuid.UUID `json:"id"`
	TimeLimit     int8
	MaxNumPlayers int32 `json:"max_num_players"`
}

func (r *Room) Validate() error {
	if r.TimeLimit == 0 {
		r.TimeLimit = 10
	}
	if r.MaxNumPlayers == 0 {
		r.MaxNumPlayers = 4
	}
	return nil
}
