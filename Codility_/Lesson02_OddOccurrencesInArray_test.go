package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__Lesson02__OddOccurrencesInArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"TestCase01", args{A: []int{9, 3, 9, 3, 9, 7, 9}}, 7},
		{"TestCase02", args{A: []int{}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, OddOccurrencesInArray(tt.args.A))
		})
	}
}

// https://app.codility.com/demo/results/training4D6HVK-BYG/
