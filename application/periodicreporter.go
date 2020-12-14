package application

import (
	"context"
	"github.com/jmwri/luckydice/domain/stats"
	"go.uber.org/zap"
	"time"
)

func NewPeriodicReporter(logger *zap.Logger, period time.Duration, guildCountProvider stats.GuildCountProvider, periodStatsProvider stats.PeriodStatsProvider) *PeriodicReporter {
	return &PeriodicReporter{
		logger:              logger,
		period:              period,
		guildCountProvider:  guildCountProvider,
		periodStatsProvider: periodStatsProvider,
	}
}

type PeriodicReporter struct {
	logger              *zap.Logger
	period              time.Duration
	guildCountProvider  stats.GuildCountProvider
	periodStatsProvider stats.PeriodStatsProvider
}

func (r *PeriodicReporter) Start(ctx context.Context) {
	ticker := time.NewTicker(r.period)
	for {
		r.report()

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			continue
		}
	}
}

func (r *PeriodicReporter) report() {
	connectedGuilds, err := r.guildCountProvider.GetGuildCount()
	if err != nil {
		r.logger.Error("failed to get connected guilds", zap.Error(err))
		return
	}

	periodStats := r.periodStatsProvider.LastPeriodStats()

	r.logger.Info(
		"periodic stat report",
		zap.Int("connectedGuilds", connectedGuilds),
		zap.String("period", r.period.String()),
		zap.Int("numRolls", periodStats.Rolls),
		zap.Int("numMisunderstandings", periodStats.Misunderstandings),
		zap.Int("numHelps", periodStats.Helps),
	)
}
