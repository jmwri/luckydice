package internal

import (
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/jmwri/luckydice/internal/port"
)

type Service interface {
	Handle(name, input string, outputReceiver port.OutputReceiver) error
	Stats() domain.Stats
}
