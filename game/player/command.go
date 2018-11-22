package player

const (
	CommandStop      = "transform.stop"
	CommandMoveUp    = "transform.handleMovement.up"
	CommandMoveDown  = "transform.handleMovement.down"
	CommandMoveLeft  = "transform.handleMovement.left"
	CommandMoveRight = "transform.handleMovement.right"
)

type Command string
