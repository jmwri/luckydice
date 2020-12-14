package stat

import (
	"github.com/jmwri/luckydice/domain"
	"github.com/jmwri/luckydice/domain/stats"
)

func NewPeriodStatsProvider(inputRecorder domain.InputRecorder) *PeriodStatsProvider {
	return &PeriodStatsProvider{
		inputRecorder: inputRecorder,
	}
}

type PeriodStatsProvider struct {
	inputRecorder domain.InputRecorder
}

func (r *PeriodStatsProvider) LastPeriodStats() stats.PeriodStats {
	st := stats.PeriodStats{
		Rolls:             r.inputRecorder.Rolls(),
		Helps:             r.inputRecorder.Helps(),
		Misunderstandings: r.inputRecorder.Misunderstandings(),
	}
	r.inputRecorder.Reset()
	return st
}
