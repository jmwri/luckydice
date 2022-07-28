package core

import (
	"github.com/jmwri/luckydice/internal/domain"
	"math/rand"
)

func Roll(input domain.RollInput) (domain.RollOutput, error) {
	rolls := make([]int, input.NumRolls)
	total := 0
	for i := 0; i < input.NumRolls; i++ {
		roll := rand.Intn(input.MaxRoll) + 1
		rolls[i] = roll
		total += roll
	}

	return domain.NewRollOutput(rolls, input.Modifier, total+input.Modifier), nil
}
