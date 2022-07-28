package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"strconv"
	"strings"
)

func GetInvalidOutput(name, messagePrefix string) string {
	return fmt.Sprintf("Sorry %s, I don't understand. You can ask me for help with `%s help`.", name, messagePrefix)
}

func GetHelpOutput(name, messagePrefix string) string {
	lines := []string{
		fmt.Sprintf("Hi %s! You can use me by typing the following: `%s {number of rolls} d{sides on die} {modifier}`. For example: `%s 2 d20 +3`.", name, messagePrefix, messagePrefix),
		fmt.Sprintf("All whitespace is optional, and you can exclude {number of rolls} and {modifier}."),
		fmt.Sprintf("`%s 1d20+0` is the same as `%s d20`.", messagePrefix, messagePrefix),
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
