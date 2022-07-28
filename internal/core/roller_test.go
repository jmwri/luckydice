package core_test

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoll(t *testing.T) {
	tests := []domain.RollInput{
		domain.NewRollInput(1, 20, 1),
		domain.NewRollInput(5, 4, 6),
		domain.NewRollInput(1, 2, -6),
	}

	for _, input := range tests {
		test := input
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			output, err := core.Roll(input)
			assert.Nil(t, err)
			assert.Equal(t, input.Modifier, output.Modifier)
			assert.Equal(t, input.NumRolls, len(output.Rolls))
			expectedResult := 0
			for _, roll := range output.Rolls {
				expectedResult += roll
			}
			expectedResult += input.Modifier
			assert.Equal(t, expectedResult, output.Result)
		})
	}
}

func TestRoll_MaxNumRolls(t *testing.T) {
	input := domain.NewRollInput(2001, 2, -6)
	_, err := core.Roll(input)
	assert.NotNil(t, err)
}

func TestRoll_MaxRoll(t *testing.T) {
	input := domain.NewRollInput(10, 1001, -6)
	_, err := core.Roll(input)
	assert.NotNil(t, err)
}
