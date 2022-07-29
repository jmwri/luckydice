package core

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/domain"
	"github.com/jmwri/luckydice/internal/port"
	"strings"
	"time"
)

func NewService(opts domain.ServiceOpts, stats port.StatsRegistry) *Service {
	return &Service{
		opts:  opts,
		stats: stats,
	}
}

type Service struct {
	opts  domain.ServiceOpts
	stats port.StatsRegistry
}

func (s *Service) HandleRoll(name, input string) (string, error) {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	rollInput, rollOutput, err := s.roll(input)
	if err != nil {
		return s.handleInvalid(name)
	}
	return s.handleSuccess(name, rollInput, rollOutput)
}

func (s *Service) handleSuccess(name string, input domain.RollInput, output domain.RollOutput) (string, error) {
	s.stats.AddRoll()
	return GetSuccessfulOutput(name, input, output), nil
}

func (s *Service) handleInvalid(name string) (string, error) {
	s.stats.AddInvalid()
	return GetInvalidOutput(name, s.opts.RollUtilCmdName, s.opts.RollUtilHelpCmdName), nil
}

func (s *Service) HandleHelp(name string) (string, error) {
	s.stats.AddHelp()
	return GetHelpOutput(name, s.opts.RollUtilCmdName, s.opts.RollUtilHelpCmdName), nil
}

func (s *Service) HandleStats(name string) (string, error) {
	s.stats.AddStat()
	stats, err := s.stats.Get(time.Now())
	if err != nil {
		return "", fmt.Errorf("failed to get stats: %w", err)
	}
	return GetStatsOutput(name, stats), nil
}

func (s *Service) HandleRaw(name, input string) (string, error) {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)

	if !strings.HasPrefix(input, s.opts.OldPrefix) {
		return "", nil
	}

	return GetUpdatedOutput(name, s.opts.RollCmdName, s.opts.RollUtilCmdName), nil
}

func (s *Service) roll(input string) (domain.RollInput, domain.RollOutput, error) {
	rollOutput := domain.RollOutput{}
	rollInput, err := ParseRoll(input)
	if err != nil {
		return rollInput, rollOutput, fmt.Errorf("failed to parse roll: %w", err)
	}
	rollOutput, err = Roll(rollInput)
	if err != nil {
		return rollInput, rollOutput, fmt.Errorf("failed to calculate: %w", err)
	}
	return rollInput, rollOutput, err
}
