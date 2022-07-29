package core_test

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/adapter"
	"github.com/jmwri/luckydice/internal/core"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var opts = domain.ServiceOpts{
	RollCmdName:          "roll",
	RollCmdInputName:     "input",
	RollUtilCmdName:      "roll-util",
	RollUtilHelpCmdName:  "help",
	RollUtilStatsCmdName: "stats",
	OldPrefix:            "!roll",
}

func TestService_HandleRoll(t *testing.T) {
	timeProvider := adapter.NewCurrentTimeProvider()
	now := timeProvider.Now()
	guildCounter := adapter.NewGuildCountProvider()
	guildCounter.SetGuildCount(1000)
	stats := adapter.NewStatsRegistry(now, guildCounter)
	svc := core.NewService(opts, stats, timeProvider)

	rand.Seed(123123)

	type testCase struct {
		name, input, expected string
	}
	tests := []testCase{
		{
			name:     "John",
			input:    "2d20+5",
			expected: "John rolled 2d20+5 [1,19]+5. Result: **25**",
		},
		{
			name:     "Jim",
			input:    "5 d20-5",
			expected: "Jim rolled 5d20-5 [6,9,19,11,6]-5. Result: **46**",
		},
		{
			name:     "Holly",
			input:    "d8",
			expected: "Holly rolled 1d8+0 [6]+0. Result: **6**",
		},
		{
			name:     "@1239876412334234",
			input:    "9d0+5",
			expected: "Sorry @1239876412334234, I don't understand. You can ask me for help with `/roll-util help`.",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s %s", test.name, test.input), func(t *testing.T) {
			actual, err := svc.HandleRoll(test.name, test.input)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestService_HandleHelp(t *testing.T) {
	timeProvider := adapter.NewCurrentTimeProvider()
	now := timeProvider.Now()
	guildCounter := adapter.NewGuildCountProvider()
	guildCounter.SetGuildCount(1000)
	stats := adapter.NewStatsRegistry(now, guildCounter)
	svc := core.NewService(opts, stats, timeProvider)

	expected := "Hi Joe! You can use me by typing the following: `/roll {number of rolls} d{sides on die} {modifier}`. For example: `/roll 2 d20 +3`." +
		"\nAll whitespace is optional, and you can exclude {number of rolls} and {modifier}." +
		"\n`/roll 1d20+0` is the same as `/roll d20`."
	actual, err := svc.HandleHelp("Joe")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestService_HandleStats(t *testing.T) {
	currentTimeProvider := adapter.NewCurrentTimeProvider()
	timeProvider := adapter.NewMockTimeProvider()
	now := currentTimeProvider.Now()
	statsCalledAt := now.Add(time.Minute * 60)

	guildCounter := adapter.NewGuildCountProvider()
	guildCounter.SetGuildCount(1000)
	stats := adapter.NewStatsRegistry(now, guildCounter)
	timeProvider.Add(statsCalledAt)
	svc := core.NewService(opts, stats, timeProvider)

	expected := "Joe, here are stats over the past 1h0m0s" +
		"\nNumber of servers: 1000" +
		"\nNumber of rolls: 0" +
		"\nNumber of invalid rolls: 0" +
		"\nNumber of help requests: 0" +
		"\nNumber of stats requests: 1" +
		"\nNumber of requests to old command: 0"
	actual, err := svc.HandleStats("Joe")
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestService_HandleRaw(t *testing.T) {
	timeProvider := adapter.NewCurrentTimeProvider()
	guildCounter := adapter.NewGuildCountProvider()
	stats := adapter.NewStatsRegistry(timeProvider.Now(), guildCounter)
	svc := core.NewService(opts, stats, timeProvider)

	type testCase struct {
		input    string
		expected string
	}
	tests := []testCase{
		{
			input:    "some random discord message",
			expected: "",
		},
		{
			input: "!roll d20",
			expected: "John, this bot is now using slash commands as suggested by discord." +
				"\nPlease use `/roll` and `/roll-util` from now on!",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			actual, err := svc.HandleRaw("John", test.input)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
