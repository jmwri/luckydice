package port

type GuildCountProvider interface {
	GetGuildCount() (int, error)
}
