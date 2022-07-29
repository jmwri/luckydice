package adapter

import (
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/jmwri/luckydice/internal/port"
	"sync"
	"time"
)

func NewStatsRegistry(start time.Time, guildCountProvider port.GuildCountProvider) *StatsRegistry {
	return &StatsRegistry{
		start:              start,
		mu:                 sync.Mutex{},
		stats:              domain.NewStats(),
		guildCountProvider: guildCountProvider,
	}
}

type StatsRegistry struct {
	start              time.Time
	mu                 sync.Mutex
	stats              domain.Stats
	guildCountProvider port.GuildCountProvider
}

func (r *StatsRegistry) AddRoll() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stats = r.stats.AddRolls(1)
}

func (r *StatsRegistry) AddHelp() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stats = r.stats.AddHelps(1)
}

func (r *StatsRegistry) AddInvalid() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stats = r.stats.AddInvalids(1)
}

func (r *StatsRegistry) AddStat() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.stats = r.stats.AddStats(1)
}

func (r *StatsRegistry) Get(now time.Time) (domain.StatsResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return domain.NewStatsResult(now.Sub(r.start), r.guildCountProvider.GetGuildCount(), r.stats.NumRoll, r.stats.NumHelp, r.stats.NumInvalid, r.stats.NumStat), nil
}
