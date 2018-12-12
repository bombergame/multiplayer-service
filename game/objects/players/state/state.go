package playerstate

type State int8

const (
	Alive = 1
	Dead  = 0
)

func (st State) ToString() string {
	switch st {
	case Alive:
		return "alive"
	case Dead:
		return "dead"
	default:
		return "unknown"
	}
}
