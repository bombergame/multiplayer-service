package rooms

type GameState int8

const (
	GameStatePending = iota
	GameStateOn
	GameStatePaused
	GameStateOff
)

func (st GameState) ToString() string {
	switch int8(st) {
	case GameStatePending:
		return "state.pending"
	case GameStateOn:
		return "state.on"
	case GameStateOff:
		return "state.off"
	default:
		return "state.unknown"
	}
}
