package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"regexp"
	"strconv"
)

func rollPattern() *regexp.Regexp {
	// Parts to make up:
	// 2 d20 +3
	patternParts := []string{
		`^`,                 // start of string
		`(?P<num>\d{0,2})?`, // 2
		`d(?P<max>\d+)`,     // d20
		`(?P<mod>[+-]\d+)?`, // +3
		`$`,
	}

	pattern := ""
	for _, part := range patternParts {
		pattern = pattern + part + `\s*`
	}
	return regexp.MustCompile(pattern)
}

func rollInputToMap(s string) map[string]string {
	pattern := rollPattern()
	match := pattern.FindStringSubmatch(s)

	paramsMap := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func ParseRoll(s string) (domain.RollInput, error) {
	mapped := rollInputToMap(s)

	roll := domain.RollInput{
		NumRolls: 1,
		MaxRoll:  0,
		Modifier: 0,
	}

	if val, ok := mapped["num"]; ok && val != "" {
		numRolls, err := strconv.Atoi(val)
		if err != nil {
			return roll, fmt.Errorf("failed to parse roll num: %w", err)
		}
		roll.NumRolls = numRolls
	}

	if val, ok := mapped["max"]; ok && val != "" {
		maxRoll, err := strconv.Atoi(val)
		if err != nil {
			return roll, fmt.Errorf("failed to parse roll max: %w", err)
		}
		roll.MaxRoll = maxRoll
	}

	if val, ok := mapped["mod"]; ok && val != "" {
		modifier, err := strconv.Atoi(val)
		if err != nil {
			return roll, fmt.Errorf("failed to parse roll mod: %w", err)
		}
		roll.Modifier = modifier
	}

	if roll.NumRolls < 1 {
		return roll, fmt.Errorf("roll num must be positive")
	}
	if roll.MaxRoll < 1 {
		return roll, fmt.Errorf("roll max must be positive")
	}

	return roll, nil
}
