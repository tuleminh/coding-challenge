package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__Lesson03_PermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"TestCase01", args{[]int{2, 3, 1, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, PermMissingElem(tt.args.A))
		})
	}
}
