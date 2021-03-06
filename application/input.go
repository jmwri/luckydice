package application

import (
	"github.com/jmwri/luckydice/domain"
	"regexp"
	"strconv"
	"sync"
)

func NewInputParser() *InputParser {
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
	return &InputParser{
		pattern: regexp.MustCompile(pattern),
	}
}

type InputParser struct {
	pattern *regexp.Regexp
}

func (r *InputParser) toMap(s string) map[string]string {
	match := r.pattern.FindStringSubmatch(s)

	paramsMap := make(map[string]string)
	for i, name := range r.pattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

func (r *InputParser) Parse(s string) (domain.RollInput, error) {
	mapped := r.toMap(s)

	roll := domain.RollInput{
		NumRolls: 1,
		MaxRoll:  0,
		Modifier: 0,
	}

	if val, ok := mapped["num"]; ok && val != "" {
		numRolls, err := strconv.Atoi(val)
		if err != nil {
			return roll, domain.ErrInvalidInput
		}
		roll.NumRolls = numRolls
	}

	if val, ok := mapped["max"]; ok && val != "" {
		maxRoll, err := strconv.Atoi(val)
		if err != nil {
			return roll, domain.ErrInvalidInput
		}
		roll.MaxRoll = maxRoll
	}

	if val, ok := mapped["mod"]; ok && val != "" {
		modifier, err := strconv.Atoi(val)
		if err != nil {
			return roll, domain.ErrInvalidInput
		}
		roll.Modifier = modifier
	}

	if roll.NumRolls < 1 {
		return roll, domain.ErrInvalidInput
	}
	if roll.MaxRoll < 1 {
		return roll, domain.ErrInvalidInput
	}

	return roll, nil
}

func NewInputRecorder() *InputRecorder {
	return &InputRecorder{
		numRolls:             0,
		numMisunderstandings: 0,
		numHelps:             0,
		mutex:                sync.Mutex{},
	}
}

type InputRecorder struct {
	numRolls             int
	numMisunderstandings int
	numHelps             int
	mutex                sync.Mutex
}

func (r *InputRecorder) RecordRoll(input domain.RollInput) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.numRolls++
}

func (r *InputRecorder) Rolls() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.numRolls
}

func (r *InputRecorder) RecordMisunderstanding() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.numMisunderstandings++
}

func (r *InputRecorder) Misunderstandings() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.numMisunderstandings
}

func (r *InputRecorder) RecordHelp() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.numHelps++
}

func (r *InputRecorder) Helps() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.numHelps
}

func (r *InputRecorder) Reset() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.numRolls = 0
	r.numMisunderstandings = 0
	r.numHelps = 0
}
