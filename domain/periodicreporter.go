package domain

import "context"

type PeriodicReporter interface {
	Start(ctx context.Context)
}
