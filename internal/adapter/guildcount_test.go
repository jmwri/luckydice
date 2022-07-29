package adapter_test

import (
	"fmt"
	"github.com/jmwri/luckydice/internal/adapter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGuildCountProvider(t *testing.T) {
	p := adapter.NewGuildCountProvider()
	assert.Equal(t, 0, p.GetGuildCount())
}

func TestGuildCountProvider_GetGuildCount(t *testing.T) {
	p := adapter.NewGuildCountProvider()
	tests := []int{1, 5, 7, 10, 38281}
	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%d", test), func(t *testing.T) {
			p.SetGuildCount(test)
			assert.Equal(t, test, p.GetGuildCount())
		})
	}
}
