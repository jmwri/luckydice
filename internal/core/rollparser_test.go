package core_test

import (
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRoll(t *testing.T) {
	type testCase struct {
		input    string
		expected domain.RollInput
		err      error
	}
	tests := []testCase{
		{
			input:    "1 d20 +5",
			expected: domain.NewRollInput(1, 20, 5),
			err:      nil,
		},
		{
			input:    "1d20+5",
			expected: domain.NewRollInput(1, 20, 5),
			err:      nil,
		},
		{
			input:    "1 d20",
			expected: domain.NewRollInput(1, 20, 0),
			err:      nil,
		},
		{
			input:    "d20",
			expected: domain.NewRollInput(1, 20, 0),
			err:      nil,
		},
		{
			input:    "d20-5",
			expected: domain.NewRollInput(1, 20, -5),
			err:      nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			actual, err := core.ParseRoll(test.input)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
