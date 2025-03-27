package cast

import (
	"errors"
	"fmt"
)

var (
	errNil      = errors.New("value is nil")
	errType     = errors.New("cannot convert value to the specified type")
	errOverflow = errors.New("value exceeds the allowable range")
)

func newNilError(typ string) error {
	return fmt.Errorf("%s: %w", typ, errNil)
}

func newTypeError(typ string) error {
	return fmt.Errorf("%s: %w", typ, errType)
}

func newOverflowError(typ string) error {
	return fmt.Errorf("%s: %w", typ, errOverflow)
}

// IsNilError returns true if the error is not nil and represents a nil value error.
func IsNilError(err error) bool {
	return errors.Is(err, errNil)
}

// IsCastError returns true if the error is not nil and represents a type casting error.
func IsCastError(err error) bool {
	return errors.Is(err, errType)
}

// IsOverflowError returns true if the error is not nil and represents a value overflow error.
func IsOverflowError(err error) bool {
	return errors.Is(err, errOverflow)
}
