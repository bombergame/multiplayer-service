package bombstate

type State string

const (
	Placed    = State("placed")
	Detonated = State("detonated")
)

func (st State) String() string {
	return string(st)
}
