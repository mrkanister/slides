package test

import (
	"testing"
)

func Test_myFunc(t *testing.T) {
	tests := map[string]struct {
		arg  int
		want int
	}{
		"test 1": {
			123, 456,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// ...
		})
	}
}
