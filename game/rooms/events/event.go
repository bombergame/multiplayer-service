package roomevents

type Event string

func (e Event) String() string {
	return string(e)
}

const (
	PlayerPrefix       = "player."
	PlayerConnected    = "connected"
	PlayerDisconnected = "disconnected"
)

const (
	GamePrefix  = "game."
	GameStarted = "started"
	GamePaused  = "paused"
	GameEnded   = "ended"
)
