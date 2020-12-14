package domain

import "errors"

var ErrInvalidInput = errors.New("invalid input")

type InputParser interface {
	Parse(s string) (RollInput, error)
}

type InputRecorder interface {
	RecordRoll(input RollInput)
	Rolls() int
	RecordMisunderstanding()
	Misunderstandings() int
	RecordHelp()
	Helps() int
	Reset()
}
