package June_LeetCoding_Challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
func Test__W1D03_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{"TestCase02", args{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}}, 110},
		{"TestCase01", args{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}}, 1859},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, twoCitySchedCost(tt.args.costs))
		})
	}
}
