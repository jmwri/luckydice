package application

import (
	"github.com/jmwri/luckydice/domain"
	"math/rand"
)

func NewRoller() *Roller {
	return &Roller{}
}

type Roller struct {
}

func (r *Roller) Roll(input domain.RollInput) domain.RollOutput {
	rolls := make([]int, input.NumRolls)
	total := 0
	for i := 0; i < input.NumRolls; i++ {
		roll := rand.Intn(input.MaxRoll-1) + 1
		rolls[i] = roll
		total += roll
	}

	output := domain.RollOutput{
		Rolls:    rolls,
		Modifier: input.Modifier,
		Result:   total + input.Modifier,
	}
	return output
}
