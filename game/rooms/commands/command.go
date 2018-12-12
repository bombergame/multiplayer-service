package gamecommands

type Cmd string

const (
	Prefix = "game."

	Start = Prefix + "start"
	Stop  = Prefix + "stop"
	End   = Prefix + "end"
)

type CmdChan chan Cmd

const (
	ChanLen = 10
)
