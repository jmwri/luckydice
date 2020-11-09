package domain

type InputParser interface {
	Parse(s string) RollInput
}
