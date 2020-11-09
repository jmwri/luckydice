package domain

type OutputBuilder interface {
	Build(name string, output RollOutput) string
}
