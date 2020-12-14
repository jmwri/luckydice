package stats

type GuildCountProvider interface {
	GetGuildCount() (int, error)
}
