package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__Lesson03__FrogJmp(t *testing.T) {
	type args struct {
		X int
		Y int
		D int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"TestCase01", args{X: 10, Y: 85, D: 30}, 3},
		{"TestCase02", args{X: 1, Y: 5, D: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FrogJmp(tt.args.X, tt.args.Y, tt.args.D))
		})
	}
}
