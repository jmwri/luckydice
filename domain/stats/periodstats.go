package stats

type PeriodStatsProvider interface {
	LastPeriodStats() PeriodStats
}

type PeriodStats struct {
	Rolls             int
	Helps             int
	Misunderstandings int
}
