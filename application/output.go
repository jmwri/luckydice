package application

import (
	"fmt"
	"github.com/jmwri/luckydice/domain"
	"strconv"
	"strings"
)

func NewOutputBuilder() *OutputBuilder {
	return &OutputBuilder{}
}

type OutputBuilder struct {
}

func (b *OutputBuilder) Build(name string, output domain.RollOutput) string {
	stringRolls := make([]string, len(output.Rolls))
	for k, v := range output.Rolls {
		stringRolls[k] = strconv.Itoa(v)
	}
	rolls := strings.Join(stringRolls, ",")
	return fmt.Sprintf("%s rolled [%s]%+d. Result: **%d**", name, rolls, output.Modifier, output.Result)
}
