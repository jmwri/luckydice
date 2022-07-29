package port

import (
	"github.com/jmwri/luckydice/internal/domain"
	"time"
)

type StatsRegistry interface {
	AddRoll()
	AddHelp()
	AddInvalid()
	AddStat()
	AddOld()
	Get(now time.Time) (domain.StatsResult, error)
}
