package commands

type Command int8

const (
	Stop      = 0
	MoveUp    = 1
	MoveDown  = 2
	MoveLeft  = 3
	MoveRight = 4
)

type Chan chan Command
