package types

// KnownError represents a known error in the API.
type KnownError struct {
	msg    string
	status int
}

func NewKnownError(msg string, status int) KnownError {
	return KnownError{
		msg:    msg,
		status: status,
	}
}

// Error implements the error interface.
func (ke KnownError) Error() string {
	return ke.msg
}

// StatusCode returns the known errors' intended HTTP status code.
func (ke KnownError) StatusCode() int {
	return ke.status
}
