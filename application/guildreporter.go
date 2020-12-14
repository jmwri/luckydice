package application

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func NewGuildReporter(dg *discordgo.Session, period time.Duration) *GuildReporter {
	return &GuildReporter{
		dg:     dg,
		period: period,
	}
}

type GuildReporter struct {
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
	pageSize := 1

	lastGuildID := ""
	totGuilds := 0
	numRequests := 0
	lastRequestSize := 0
	for numRequests == 0 || lastRequestSize >= pageSize {
		guilds, err := r.dg.UserGuilds(pageSize, "", lastGuildID)
		if err != nil {
			log.Println(err)
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
	log.Printf("Currently connected to %d guilds", totGuilds)
}
