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

func GetSuccessfulOutput(name string, output domain.RollOutput) string {
	stringRolls := make([]string, len(output.Rolls))
	for k, v := range output.Rolls {
		stringRolls[k] = strconv.Itoa(v)
	}
	rolls := strings.Join(stringRolls, ",")
	return fmt.Sprintf("%s rolled [%s]%+d. Result: **%d**", name, rolls, output.Modifier, output.Result)
}
