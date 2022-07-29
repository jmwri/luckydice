package domain

import "time"

func NewStats() Stats {
	return Stats{}
}

type Stats struct {
	NumRoll    int
	NumHelp    int
	NumInvalid int
	NumStat    int
}

func (s Stats) AddRolls(n int) Stats {
	s.NumRoll += n
	return s
}

func (s Stats) AddHelps(n int) Stats {
	s.NumHelp += n
	return s
}

func (s Stats) AddInvalids(n int) Stats {
	s.NumInvalid += n
	return s
}

func (s Stats) AddStats(n int) Stats {
	s.NumStat += n
	return s
}

func NewStatsResult(period time.Duration, numGuild, numRoll, numHelp, numInvalid, numStat int) StatsResult {
	return StatsResult{
		Period:     period,
		NumGuild:   numGuild,
		NumRoll:    numRoll,
		NumHelp:    numHelp,
		NumInvalid: numInvalid,
		NumStat:    numStat,
	}
}

type StatsResult struct {
	Period     time.Duration
	NumGuild   int
	NumRoll    int
	NumHelp    int
	NumInvalid int
	NumStat    int
}
