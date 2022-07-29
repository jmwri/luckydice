package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"strconv"
	"strings"
)

func GetInvalidOutput(name string, commands ...string) string {
	cmd := strings.Join(commands, " ")
	return fmt.Sprintf("Sorry %s, I don't understand. You can ask me for help with `/%s`.", name, cmd)
}

func GetHelpOutput(name string, commands ...string) string {
	cmd := strings.Join(commands, " ")
	lines := []string{
		fmt.Sprintf("Hi %s! You can use me by typing the following: `/%s {number of rolls} d{sides on die} {modifier}`. For example: `/%s 2 d20 +3`.", name, cmd, cmd),
		fmt.Sprintf("All whitespace is optional, and you can exclude {number of rolls} and {modifier}."),
		fmt.Sprintf("`/%s 1d20+0` is the same as `/%s d20`.", cmd, cmd),
	}
	return strings.Join(lines, "\n")
}

func GetStatsOutput(name string, stats domain.StatsResult) string {
	lines := []string{
		fmt.Sprintf("%s, here are stats over the past %s", name, stats.Period.String()),
		fmt.Sprintf("Number of servers: %d", stats.NumGuild),
		fmt.Sprintf("Number of rolls: %d", stats.NumRoll),
		fmt.Sprintf("Number of invalid rolls: %d", stats.NumInvalid),
		fmt.Sprintf("Number of help requests: %d", stats.NumHelp),
		fmt.Sprintf("Number of stats requests: %d", stats.NumStat),
	}
	return strings.Join(lines, "\n")
}

func GetSuccessfulOutput(name string, input domain.RollInput, output domain.RollOutput) string {
	modSymbol := "+"
	if input.Modifier < 0 {
		modSymbol = ""
	}
	inputStr := fmt.Sprintf("%dd%d%s%d", input.NumRolls, input.MaxRoll, modSymbol, input.Modifier)

	stringRolls := make([]string, len(output.Rolls))
	for k, v := range output.Rolls {
		stringRolls[k] = strconv.Itoa(v)
	}
	rolls := strings.Join(stringRolls, ",")
	return fmt.Sprintf("%s rolled (%s) [%s]%+d. Result: **%d**", name, inputStr, rolls, output.Modifier, output.Result)
}

func GetUpdatedOutput(name, rollCmdName, rollUtilCmdName string) string {
	lines := []string{
		fmt.Sprintf("%s, this bot is now using slash commands as suggested by discord.", name),
		fmt.Sprintf("Please use `/%s` and `/%s` from now on!", rollCmdName, rollUtilCmdName),
	}
	return strings.Join(lines, "\n")
}
