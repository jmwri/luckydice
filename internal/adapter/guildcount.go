package adapter

import (
	"github.com/bwmarrin/discordgo"
)

func NewGuildCountProvider(dg *discordgo.Session) *GuildCountProvider {
	return &GuildCountProvider{
		dg: dg,
	}
}

type GuildCountProvider struct {
	dg *discordgo.Session
}

func (r *GuildCountProvider) GetGuildCount() (int, error) {
	pageSize := 100

	lastGuildID := ""
	totGuilds := 0
	numRequests := 0
	lastRequestSize := 0
	for numRequests == 0 || lastRequestSize >= pageSize {
		guilds, err := r.dg.UserGuilds(pageSize, "", lastGuildID)
		if err != nil {
			return totGuilds, err
		}
		numRequests++
		totGuilds += len(guilds)
		lastRequestSize = len(guilds)
		if lastRequestSize == 0 {
			continue
		}
		lastGuildID = guilds[lastRequestSize-1].ID
	}
	return totGuilds, nil
}
