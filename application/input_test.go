package application_test

import (
	"errors"
	"github.com/jmwri/luckydice/application"
	"github.com/jmwri/luckydice/domain"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	t.Parallel()
	parser := application.NewInputParser()

	tests := []struct {
		name  string
		input string
		exp   domain.RollInput
	}{
		{"full with positive mod", "2d6+2", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  6,
			Modifier: 2,
		}},
		{"full with negative mod", "2d6-2", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  6,
			Modifier: -2,
		}},
		{"no num with positive mod", "d6+3", domain.RollInput{
			NumRolls: 1,
			MaxRoll:  6,
			Modifier: 3,
		}},
		{"no num with negative mod", "d6-3", domain.RollInput{
			NumRolls: 1,
			MaxRoll:  6,
			Modifier: -3,
		}},
		{"only max roll", "d6", domain.RollInput{
			NumRolls: 1,
			MaxRoll:  6,
			Modifier: 0,
		}},
		{"num without modifier", "2d12", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  12,
			Modifier: 0,
		}},
		{"full with positive mod - whitespace", "2 d6 +2", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  6,
			Modifier: 2,
		}},
		{"full with negative mod - whitespace", "2 d6 -2", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  6,
			Modifier: -2,
		}},
		{"no num with positive mod - whitespace", "d6 +3", domain.RollInput{
			NumRolls: 1,
			MaxRoll:  6,
			Modifier: 3,
		}},
		{"no num with negative mod - whitespace", "d6 -3", domain.RollInput{
			NumRolls: 1,
			MaxRoll:  6,
			Modifier: -3,
		}},
		{"num without modifier - whitespace", "2 d12", domain.RollInput{
			NumRolls: 2,
			MaxRoll:  12,
			Modifier: 0,
		}},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			out, err := parser.Parse(test.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			} else if out != test.exp {
				t.Errorf("expected %v, got %v", test.exp, out)
			}
		})
	}
}

func TestParser_Parse_InvalidInput(t *testing.T) {
	t.Parallel()
	parser := application.NewInputParser()

	tests := []struct {
		name  string
		input string
	}{
		{"num rolls must be at least 1", "0d6"},
		{"max roll must be at least 1", "1d0"},
		{"invalid pattern", "xd5-1"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := parser.Parse(test.input)
			if !errors.Is(err, domain.ErrInvalidInput) {
				t.Errorf("expected %s, got %s", domain.ErrInvalidInput, err)
			}
		})
	}
}
