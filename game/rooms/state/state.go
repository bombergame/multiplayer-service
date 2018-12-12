package gamestate

type State int8

const (
	Pending = iota
	On
	Paused
	Off
)

func (st State) ToString() string {
	switch int8(st) {
	case Pending:
		return "pending"
	case Paused:
		return "paused"
	case On:
		return "on"
	case Off:
		return "off"
	default:
		return "unknown"
	}
}
