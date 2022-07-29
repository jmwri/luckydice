package domain_test

import (
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewStats(t *testing.T) {
	s := domain.NewStats()
	assert.Equal(t, 0, s.NumRoll)
	assert.Equal(t, 0, s.NumHelp)
	assert.Equal(t, 0, s.NumInvalid)
	assert.Equal(t, 0, s.NumStat)
}

func TestStats_AddRolls(t *testing.T) {
	s := domain.NewStats()
	s = s.AddRolls(5)
	assert.Equal(t, 5, s.NumRoll)
}

func TestStats_AddHelps(t *testing.T) {
	s := domain.NewStats()
	s = s.AddHelps(5)
	assert.Equal(t, 5, s.NumHelp)
}

func TestStats_AddInvalids(t *testing.T) {
	s := domain.NewStats()
	s = s.AddInvalids(5)
	assert.Equal(t, 5, s.NumInvalid)
}

func TestNewStatsResult(t *testing.T) {
	period := time.Minute * 50
	guilds := 200
	rolls := 100
	helps := 80
	invalids := 50
	stats := 30
	r := domain.NewStatsResult(period, guilds, rolls, helps, invalids, stats)
	assert.Equal(t, period, r.Period)
	assert.Equal(t, guilds, r.NumGuild)
	assert.Equal(t, rolls, r.NumRoll)
	assert.Equal(t, helps, r.NumHelp)
	assert.Equal(t, invalids, r.NumInvalid)
	assert.Equal(t, stats, r.NumStat)
}
