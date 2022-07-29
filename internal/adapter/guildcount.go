package adapter

func NewGuildCountProvider() *GuildCountProvider {
	return &GuildCountProvider{
		count: 0,
	}
}

type GuildCountProvider struct {
	count int
}

func (r *GuildCountProvider) GetGuildCount() int {
	return r.count
}

func (r *GuildCountProvider) SetGuildCount(num int) {
	r.count = num
}
