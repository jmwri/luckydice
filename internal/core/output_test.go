package core_test

import (
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInvalidOutput(t *testing.T) {
	output := core.GetInvalidOutput("John", "!roll")
	expected := "Sorry John, I don't understand. You can ask me for help with `!roll help`."
	assert.Equal(t, expected, output)
}

func TestGetHelpOutput(t *testing.T) {
	output := core.GetHelpOutput("John", "!roll")
	expected := "Hi John! You can use me by typing the following: `!roll {number of rolls} d{sides on die} {modifier}`. For example: `!roll 2 d20 +3`." +
		"\nAll whitespace is optional, and you can exclude {number of rolls} and {modifier}." +
		"\n`!roll 1d20+0` is the same as `!roll d20`."
	assert.Equal(t, expected, output)
}

func TestGetSuccessfulOutput(t *testing.T) {
	rollOutput := domain.NewRollOutput([]int{1, 2, 3}, 2, 8)
	output := core.GetSuccessfulOutput("John", rollOutput)
	expected := "John rolled [1,2,3]+2. Result: **8**"
	assert.Equal(t, expected, output)
}
