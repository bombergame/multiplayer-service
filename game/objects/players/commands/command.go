package playercommands

type Cmd string

const (
	Prefix = "player."

	MovementPrefix = "move."
	MoveUp         = MovementPrefix + "up"
	MoveDown       = MovementPrefix + "down"
	MoveLeft       = MovementPrefix + "left"
	MoveRight      = MovementPrefix + "right"

	DropBomb = Prefix + "drop.bomb"
)

type CmdChan chan Cmd

const (
	ChanLen = 10
)
