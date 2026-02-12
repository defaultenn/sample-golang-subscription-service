package constants

import "errors"

var (
	ErrIncorrectInputData  = errors.New("incorrect request data")
	ErrInternalServerError = errors.New("internal server error")
)
