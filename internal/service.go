package internal

import (
	"github.com/jmwri/luckydice/internal/domain"
)

type Service interface {
	Handle(name, input string) (string, error)
	Stats() domain.Stats
}
