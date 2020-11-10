package domain

import "errors"

var ErrInvalidInput = errors.New("invalid input")

type InputParser interface {
	Parse(s string) (RollInput, error)
}
