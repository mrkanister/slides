package test

import (
	"testing"
)

func Test_myFunc(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"test 1",
			123, 456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ...
		})
	}
}
