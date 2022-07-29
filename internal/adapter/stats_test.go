package adapter_test

import (
	"github.com/jmwri/luckydice/internal/adapter"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewStatsRegistry_Get(t *testing.T) {
	now := time.Now()
	after := now.Add(time.Minute * 10)
	guildCountProvider := adapter.NewGuildCountProvider()
	guildCountProvider.SetGuildCount(1000)
	r := adapter.NewStatsRegistry(now, guildCountProvider)
	for i := 0; i < 5; i++ {
		r.AddRoll()
	}
	for i := 0; i < 10; i++ {
		r.AddHelp()
	}
	for i := 0; i < 15; i++ {
		r.AddInvalid()
	}
	for i := 0; i < 20; i++ {
		r.AddStat()
	}
	actual, err := r.Get(after)
	expected := domain.NewStatsResult(time.Minute*10, 1000, 5, 10, 15, 20)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
