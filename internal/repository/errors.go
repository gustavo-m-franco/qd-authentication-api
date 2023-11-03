package repository

// Error is the error type for the repository
type Error struct {
	Message string
}

// Error returns the error message
func (e *Error) Error() string {
	return e.Message
}
