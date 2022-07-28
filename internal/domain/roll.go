package domain

func NewRollInput(rolls, max, mod int) RollInput {
	return RollInput{
		NumRolls: rolls,
		MaxRoll:  max,
		Modifier: mod,
	}
}

type RollInput struct {
	NumRolls int
	MaxRoll  int
	Modifier int
}

func NewRollOutput(rolls []int, modifier, result int) RollOutput {
	return RollOutput{
		Rolls:    rolls,
		Modifier: modifier,
		Result:   result,
	}
}

type RollOutput struct {
	Rolls    []int
	Modifier int
	Result   int
}
