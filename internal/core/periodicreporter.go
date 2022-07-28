package core

import (
	"context"
	"github.com/jmwri/luckydice/internal"
	"github.com/jmwri/luckydice/internal/port"
	"go.uber.org/zap"
	"time"
)

func NewPeriodicReporter(logger *zap.Logger, period time.Duration, guildCountProvider port.GuildCountProvider, svc internal.Service) *PeriodicReporter {
	return &PeriodicReporter{
		logger:             logger,
		period:             period,
		guildCountProvider: guildCountProvider,
		svc:                svc,
	}
}

type PeriodicReporter struct {
	logger             *zap.Logger
	period             time.Duration
	guildCountProvider port.GuildCountProvider
	svc                internal.Service
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

	periodStats := r.svc.Stats()

	r.logger.Info(
		"periodic stat report",
		zap.Int("connectedGuilds", connectedGuilds),
		zap.String("period", r.period.String()),
		zap.Int64("numRolls", periodStats.NumRoll),
		zap.Int64("numInvalid", periodStats.NumInvalid),
		zap.Int64("numHelps", periodStats.NumHelp),
	)
}
