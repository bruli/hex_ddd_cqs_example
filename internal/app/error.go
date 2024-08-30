package app

import "fmt"

type InvalidQueryError struct {
	had, expected string
}

func NewInvalidQueryError(had, expected string) InvalidQueryError {
	return InvalidQueryError{had: had, expected: expected}
}

func (i InvalidQueryError) Error() string {
	return fmt.Sprintf("invalid query. Expected: %q, had: %q", i.expected, i.had)
}
