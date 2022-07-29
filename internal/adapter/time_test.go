package adapter_test

import (
	"github.com/jmwri/luckydice/internal/adapter"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrentTimeProvider_Now(t *testing.T) {
	p := adapter.NewCurrentTimeProvider()
	actual := p.Now()
	assert.IsType(t, time.Time{}, actual)
}

func TestMockTimeProvider_Now(t *testing.T) {
	p := adapter.NewMockTimeProvider()

	tests := []time.Time{
		time.Now(),
		time.Now().Add(time.Minute),
		time.Now().Add(time.Minute * 5),
		time.Now().Add(time.Minute * 10),
		time.Now().Add(time.Minute * 30),
		time.Now().Add(time.Minute * 1000),
	}

	p.Add(tests...)

	for _, test := range tests {
		test := test
		t.Run(test.String(), func(t *testing.T) {
			assert.Equal(t, test, p.Now())
		})
	}
}
