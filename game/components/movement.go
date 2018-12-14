package components

//go:generate easyjson

import (
	"time"
)

type Movement struct {
	MinStepInterval  time.Duration
	LastStepInterval time.Duration
}

func (m Movement) GetMessageData() MovementMessageData {
	return MovementMessageData{
		MinStepInterval:  m.MinStepInterval.Seconds(),
		LastStepInterval: m.LastStepInterval.Seconds(),
	}
}

//easyjson:json
type MovementMessageData struct {
	MinStepInterval  float64 `json:"min_step_interval"`
	LastStepInterval float64 `json:"last_step_interval"`
}
