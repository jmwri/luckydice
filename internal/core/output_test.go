package core_test

import (
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetInvalidOutput(t *testing.T) {
	output := core.GetInvalidOutput("John", "cmd", "sub", "help")
	expected := "Sorry John, I don't understand. You can ask me for help with `/cmd sub help`."
	assert.Equal(t, expected, output)
}

func TestGetHelpOutput(t *testing.T) {
	output := core.GetHelpOutput("John", "cmd", "sub", "roll")
	expected := "Hi John! You can use me by typing the following: `/cmd sub roll {number of rolls} d{sides on die} {modifier}`. For example: `/cmd sub roll 2 d20 +3`." +
		"\nAll whitespace is optional, and you can exclude {number of rolls} and {modifier}." +
		"\n`/cmd sub roll 1d20+0` is the same as `/cmd sub roll d20`."
	assert.Equal(t, expected, output)
}

func TestGetStatsOutput(t *testing.T) {
	period := time.Minute * 68
	stats := domain.NewStatsResult(period, 200, 100, 80, 60, 40, 20)
	output := core.GetStatsOutput("John", stats)
	expected := "John, here are stats over the past 1h8m0s" +
		"\nNumber of servers: 200" +
		"\nNumber of rolls: 100" +
		"\nNumber of invalid rolls: 60" +
		"\nNumber of help requests: 80" +
		"\nNumber of stats requests: 40" +
		"\nNumber of requests to old command: 20"
	assert.Equal(t, expected, output)
}

func TestGetSuccessfulOutput(t *testing.T) {
	rollInput := domain.NewRollInput(3, 20, 2)
	rollOutput := domain.NewRollOutput([]int{1, 2, 3}, 2, 8)
	output := core.GetSuccessfulOutput("John", rollInput, rollOutput)
	expected := "John rolled 3d20+2 [1,2,3]+2. Result: **8**"
	assert.Equal(t, expected, output)
}

func TestGetUpdatedOutput(t *testing.T) {
	output := core.GetUpdatedOutput("John", "roll", "roll-util")
	expected := "John, this bot is now using slash commands as suggested by discord." +
		"\nPlease use `/roll` and `/roll-util` from now on!"
	assert.Equal(t, expected, output)
}
