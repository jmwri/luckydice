package domain_test

import (
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRollInput(t *testing.T) {
	rolls := 5
	max := 20
	mod := 2
	i := domain.NewRollInput(rolls, max, mod)
	assert.Equal(t, rolls, i.NumRolls)
	assert.Equal(t, max, i.MaxRoll)
	assert.Equal(t, mod, i.Modifier)
}

func TestNewRollOutput(t *testing.T) {
	rolls := []int{2, 4, 6}
	mod := 5
	res := 17
	r := domain.NewRollOutput(rolls, mod, res)
	assert.Equal(t, rolls, r.Rolls)
	assert.Equal(t, mod, r.Modifier)
	assert.Equal(t, res, r.Result)
}
