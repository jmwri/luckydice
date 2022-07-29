package port

type GuildCountProvider interface {
	GetGuildCount() int
	SetGuildCount(num int)
}
