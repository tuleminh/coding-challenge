package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__Lesson02__CyclicRotation(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"TestCase01", args{A: []int{3, 8, 9, 7, 6}, K: 3}, []int{9, 7, 6, 3, 8}},
		{"TestCase02", args{A: []int{0, 0, 0}, K: 1}, []int{0, 0, 0}},
		{"TestCase03", args{A: []int{1, 2, 3, 4}, K: 4}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CyclicRotation(tt.args.A, tt.args.K))
		})
	}
}
