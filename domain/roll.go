package domain

type Roller interface {
	Roll(input RollInput) RollOutput
}

type RollInput struct {
	NumRolls int
	MaxRoll int
	Modifier int
}

type RollOutput struct {
	Rolls []int
	Modifier int
	Result int
}
