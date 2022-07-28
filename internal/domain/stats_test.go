package domain_test

import (
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModifiableStats_AddRoll(t *testing.T) {
	s := domain.NewStats()
	s = s.AddRoll()
	assert.Equal(t, 1, s.NumRoll)
}

func TestModifiableStats_AddRolls(t *testing.T) {
	s := domain.NewStats()
	s = s.AddRolls(5)
	assert.Equal(t, 5, s.NumRoll)
}

func TestModifiableStats_AddHelp(t *testing.T) {
	s := domain.NewStats()
	s = s.AddHelp()
	assert.Equal(t, 1, s.NumHelp)
}

func TestModifiableStats_AddHelps(t *testing.T) {
	s := domain.NewStats()
	s = s.AddHelps(5)
	assert.Equal(t, 5, s.NumHelp)
}

func TestModifiableStats_AddInvalid(t *testing.T) {
	s := domain.NewStats()
	s = s.AddInvalid()
	assert.Equal(t, 1, s.NumInvalid)
}

func TestModifiableStats_AddInvalids(t *testing.T) {
	s := domain.NewStats()
	s = s.AddInvalids(5)
	assert.Equal(t, 5, s.NumInvalid)
}

func TestModifiableStats_Reset(t *testing.T) {
	s := domain.NewStats()
	s = s.AddRoll()
	s = s.AddHelp()
	s = s.AddInvalid()
	s = s.Reset()
	assert.Equal(t, 0, s.NumRoll)
	assert.Equal(t, 0, s.NumHelp)
	assert.Equal(t, 0, s.NumInvalid)
}

func TestModifiableStats_Result(t *testing.T) {
	s := domain.NewStats()
	s = s.AddRoll()
	s = s.AddHelp()
	s = s.AddInvalid()
	res := s.Result()
	assert.Equal(t, 1, s.NumRoll)
	assert.Equal(t, 1, s.NumHelp)
	assert.Equal(t, 1, s.NumInvalid)
	assert.Equal(t, 1, res.NumRoll)
	assert.Equal(t, 1, res.NumHelp)
	assert.Equal(t, 1, res.NumInvalid)
}
