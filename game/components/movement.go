package components

import (
	"time"
)

type Movement struct {
	StepInterval     time.Duration
	LastStepInterval time.Duration
}
