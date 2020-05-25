package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__Lesson01__BinaryGap(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"TestCase01", args{N: 1041}, 5},
		{"TestCase02", args{N: 32}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, BinaryGap(tt.args.N))
		})
	}
}
