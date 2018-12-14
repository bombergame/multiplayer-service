package errs

type InvalidCellIndexError struct {
	GameError
}

func NewInvalidCellIndexError() *InvalidCellIndexError {
	return &InvalidCellIndexError{
		GameError: GameError{
			message: "cell index out of range",
		},
	}
}
