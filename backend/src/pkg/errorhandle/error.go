package errorhandle

import (
	"errors"
	"fmt"
)

//---common error handling---

func NewError(message string) error {
	return errors.New(message)
}

func Wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

func IsErrorType(err error, target error) bool {
	return errors.Is(err, target)
}

//---typical error message---

const (
	LibraryError = "library error"
)

func NewLibraryError() error {
	return NewError(LibraryError)
}
