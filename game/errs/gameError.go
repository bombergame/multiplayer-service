package errs

type GameError struct {
	message string
}

func NewGameError(message string) *GameError {
	return &GameError{
		message: message,
	}
}

func (err *GameError) Error() string {
	return err.message
}
