package June_LeetCoding_Challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name     string
		args     args
		expected []byte
	}{
		{"TestCase01", args{[]byte("hello")}, []byte("olleh")},
		{"TestCase02", args{[]byte("Hannah")}, []byte("hannaH")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			assert.Equal(t, tt.expected, tt.args.s)
		})
	}
}
