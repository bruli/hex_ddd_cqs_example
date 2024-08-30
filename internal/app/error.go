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

type InvalidCommandError struct {
	had, expected string
}

func NewInvalidCommandError(had, expected string) InvalidCommandError {
	return InvalidCommandError{had: had, expected: expected}
}

func (i InvalidCommandError) Error() string {
	return fmt.Sprintf("invalid query. Expected: %q, had: %q", i.expected, i.had)
}
