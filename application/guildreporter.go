package application

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
	"time"
)

func NewGuildReporter(logger *zap.Logger, dg *discordgo.Session, period time.Duration) *GuildReporter {
	return &GuildReporter{
		logger: logger,
		dg:     dg,
		period: period,
	}
}

type GuildReporter struct {
	logger *zap.Logger
	dg     *discordgo.Session
	period time.Duration
}

func (r *GuildReporter) Start(ctx context.Context) {
	ticker := time.NewTicker(r.period)
	for {
		r.logGuilds()
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			continue
		}
	}
}

func (r *GuildReporter) logGuilds() {
	pageSize := 100

	lastGuildID := ""
	totGuilds := 0
	numRequests := 0
	lastRequestSize := 0
	for numRequests == 0 || lastRequestSize >= pageSize {
		guilds, err := r.dg.UserGuilds(pageSize, "", lastGuildID)
		if err != nil {
			r.logger.Error("failed to request discord guilds", zap.Error(err))
			return
		}
		numRequests++
		totGuilds += len(guilds)
		lastRequestSize = len(guilds)
		if lastRequestSize == 0 {
			continue
		}
		lastGuildID = guilds[lastRequestSize-1].ID
	}
	r.logger.Info("checked guild membership", zap.Int("count", totGuilds))
}
