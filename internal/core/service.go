package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"strings"
)

func NewService(messagePrefix string) *Service {
	return &Service{
		messagePrefix: messagePrefix,
		stats:         domain.NewStats(),
	}
}

type Service struct {
	messagePrefix string
	stats         domain.ModifiableStats
}

func (s *Service) Handle(name, input string) (string, error) {
	input = strings.ToLower(input)

	if !strings.HasPrefix(input, s.messagePrefix) {
		return "", nil
	}
	input = strings.TrimPrefix(input, s.messagePrefix)
	input = strings.TrimSpace(input)

	if input == "help" {
		s.stats.AddHelp()
		return GetHelpOutput(name, s.messagePrefix), nil
	}
	output, err := s.roll(input)
	if err != nil {
		s.stats.AddInvalid()
		return GetInvalidOutput(name, s.messagePrefix), nil
	}
	s.stats.AddRoll()
	return GetSuccessfulOutput(name, output), nil
}

func (s *Service) roll(input string) (domain.RollOutput, error) {
	output := domain.RollOutput{}
	roll, err := ParseRoll(input)
	if err != nil {
		return output, fmt.Errorf("failed to parse roll: %w", err)
	}
	output, err = Roll(roll)
	if err != nil {
		return output, fmt.Errorf("failed to calculate")
	}
	return output, err
}

func (s *Service) Stats() domain.Stats {
	stats := s.stats.Result()
	s.stats = s.stats.Reset()
	return stats
}
