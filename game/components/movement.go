package components

//go:generate easyjson

import (
	"github.com/bombergame/multiplayer-service/game/physics"
	"time"
)

type Movement struct {
	MinStepInterval  time.Duration
	LastStepInterval time.Duration
	StepSize         physics.Integer
}

func (m Movement) GetMessageData() MovementMessageData {
	return MovementMessageData{
		MinStepInterval: m.MinStepInterval.Seconds(),
		StepSize:        m.StepSize,
	}
}

//easyjson:json
type MovementMessageData struct {
	StepSize        physics.Integer `json:"step_size"`
	MinStepInterval float64         `json:"min_step_interval"`
}
