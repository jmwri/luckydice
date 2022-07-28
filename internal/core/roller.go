package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"math/rand"
)

func Roll(input domain.RollInput) (domain.RollOutput, error) {
	output := domain.RollOutput{}
	if input.NumRolls > 200 {
		return output, fmt.Errorf("max number of rolls is 200")
	}
	if input.MaxRoll > 1000 {
		return output, fmt.Errorf("max roll is 1000")
	}
	rolls := make([]int, input.NumRolls)
	total := 0
	for i := 0; i < input.NumRolls; i++ {
		roll := rand.Intn(input.MaxRoll) + 1
		rolls[i] = roll
		total += roll
	}

	return domain.NewRollOutput(rolls, input.Modifier, total+input.Modifier), nil
}
