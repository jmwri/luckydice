package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"strings"
)

func NewService(opts domain.ServiceOpts) *Service {
	return &Service{
		opts:  opts,
		stats: domain.NewStats(),
	}
}

type Service struct {
	opts  domain.ServiceOpts
	stats domain.ModifiableStats
}

func (s *Service) HandleRoll(name, input string) (string, error) {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	output, err := s.roll(input)
	if err != nil {
		return s.handleInvalid(name)
	}
	return s.handleSuccess(name, output)
}

func (s *Service) handleSuccess(name string, output domain.RollOutput) (string, error) {
	s.stats.AddRoll()
	return GetSuccessfulOutput(name, output), nil
}

func (s *Service) handleInvalid(name string) (string, error) {
	s.stats.AddInvalid()
	return GetInvalidOutput(name, s.opts.RollCmdName), nil
}

func (s *Service) HandleHelp(name string) (string, error) {
	s.stats.AddHelp()
	return GetHelpOutput(name, s.opts.RollUtilCmdName, s.opts.RollUtilHelpCmdName), nil
}

func (s *Service) HandleStats(name string) (string, error) {
	s.stats.AddStat()
	return GetHelpOutput(name, s.opts.RollUtilCmdName, s.opts.RollUtilHelpCmdName), nil
}

func (s *Service) HandleRaw(name, input string) (string, error) {
	input = strings.ToLower(input)

	if !strings.HasPrefix(input, s.opts.OldPrefix) {
		return "", nil
	}
	input = strings.TrimPrefix(input, s.opts.OldPrefix)
	input = strings.TrimSpace(input)

	if input == "help" {
		return s.HandleHelp(name)
	}
	return s.HandleRoll(name, input)
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
