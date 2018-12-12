package playercommands

type Cmd string

const (
	Prefix    = "player."
	Stop      = Prefix + "stop"
	MoveUp    = Prefix + "move.up"
	MoveDown  = Prefix + "move.down"
	MoveLeft  = Prefix + "move.left"
	MoveRight = Prefix + "move.right"
)

type CmdChan chan Cmd

const (
	ChanLen = 10
)
