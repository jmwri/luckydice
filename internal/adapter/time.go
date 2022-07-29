package adapter

import "time"

func NewCurrentTimeProvider() *CurrentTimeProvider {
	return &CurrentTimeProvider{}
}

type CurrentTimeProvider struct {
}

func (p *CurrentTimeProvider) Now() time.Time {
	return time.Now().UTC()
}

func NewMockTimeProvider() *MockTimeProvider {
	return &MockTimeProvider{
		times: make([]time.Time, 0),
	}
}

type MockTimeProvider struct {
	times []time.Time
}

func (p *MockTimeProvider) Now() time.Time {
	now := p.times[0]
	if len(p.times) > 1 {
		p.times = p.times[1:]
	} else {
		p.times = make([]time.Time, 0)
	}
	return now
}

func (p *MockTimeProvider) Add(times ...time.Time) {
	p.times = append(p.times, times...)
}
