package application_test

import (
	"github.com/jmwri/luckydice/application"
	"github.com/jmwri/luckydice/domain"
	"testing"
)

func TestOutputBuilder_Build(t *testing.T) {
	t.Parallel()
	builder := application.NewOutputBuilder()

	tests := []struct {
		name        string
		discordName string
		output      domain.RollOutput
		exp         string
	}{
		{"fasdf", "Jim", domain.RollOutput{
			Rolls:    []int{4, 5, 6},
			Modifier: 2,
			Result:   17,
		}, "Jim rolled [4,5,6]+2. Result: **17**"},
		{"fasdf", "Jim", domain.RollOutput{
			Rolls:    []int{4, 5, 6},
			Modifier: -2,
			Result:   15,
		}, "Jim rolled [4,5,6]-2. Result: **15**"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			out := builder.Build(test.discordName, test.output)
			if out != test.exp {
				t.Errorf("expected %v, got %v", test.exp, out)
			}
		})
	}
}
