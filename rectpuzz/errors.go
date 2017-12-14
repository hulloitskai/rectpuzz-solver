package rectpuzz

import "fmt"

type UnexpectedZeroError struct {
	field string
}

func (err UnexpectedZeroError) Error() string {
	return fmt.Sprintf("UnexpectedZeroError: '%s' must be non-zero", err.field)
}
