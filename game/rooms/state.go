package rooms

type GameState int

const (
	GameStatePending = iota
	GameStateOn
	GameStatePaused
	GameStateOff
)
