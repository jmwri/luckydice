package application_test

import (
	"fmt"
	"github.com/jmwri/luckydice/application"
	"github.com/jmwri/luckydice/domain"
	"math/rand"
	"testing"
)

func TestRoller_Roll(t *testing.T) {
	roller := application.NewRoller()

	tests := []struct {
		name  string
		seed  int64
		input domain.RollInput
		exp   domain.RollOutput
	}{
		{"2d20+2", 123456, domain.RollInput{
			NumRolls: 2,
			MaxRoll:  20,
			Modifier: 2,
		}, domain.RollOutput{
			Rolls:    []int{13, 2},
			Modifier: 2,
			Result:   17,
		}},
		{"4d4-3", 456789, domain.RollInput{
			NumRolls: 4,
			MaxRoll:  4,
			Modifier: -3,
		}, domain.RollOutput{
			Rolls:    []int{1, 3, 3, 2},
			Modifier: -3,
			Result:   6,
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(test.seed)
			out := roller.Roll(test.input)
			outStr := fmt.Sprintf("%v", out)
			expStr := fmt.Sprintf("%v", test.exp)

			if outStr != expStr {
				t.Errorf("expected %v, got %v", test.exp, out)
			}
		})
	}
}
