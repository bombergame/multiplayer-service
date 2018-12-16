package errs

type FullQueueError struct {
	GameError
}

func NewFullQueueError() *FullQueueError {
	return &FullQueueError{
		GameError: GameError{
			message: "cache queue is full",
		},
	}
}

type EmptyQueueError struct {
	GameError
}

func NewEmptyQueueError() *EmptyQueueError {
	return &EmptyQueueError{
		GameError: GameError{
			message: "cache queue is empty",
		},
	}
}
