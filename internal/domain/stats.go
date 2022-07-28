package domain

func NewStats() ModifiableStats {
	return ModifiableStats{
		Stats: Stats{},
	}
}

type Stats struct {
	NumRoll    int64
	NumHelp    int64
	NumInvalid int64
}

type ModifiableStats struct {
	Stats
}

func (s ModifiableStats) AddRolls(n int64) ModifiableStats {
	s.NumRoll += n
	return s
}

func (s ModifiableStats) AddRoll() ModifiableStats {
	return s.AddRolls(1)
}

func (s ModifiableStats) AddHelps(n int64) ModifiableStats {
	s.NumHelp += n
	return s
}

func (s ModifiableStats) AddHelp() ModifiableStats {
	return s.AddHelps(1)
}

func (s ModifiableStats) AddInvalids(n int64) ModifiableStats {
	s.NumInvalid += n
	return s
}

func (s ModifiableStats) AddInvalid() ModifiableStats {
	return s.AddInvalids(1)
}

func (s ModifiableStats) Reset() ModifiableStats {
	return NewStats()
}

func (s ModifiableStats) Result() Stats {
	return s.Stats
}
