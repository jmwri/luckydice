package internal

type Service interface {
	HandleRoll(name, input string) (string, error)
	HandleHelp(name string) (string, error)
	HandleStats(name string) (string, error)
	HandleRaw(name, input string) (string, error)
}
