package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseRunes(c.in)
		assert.Equal(t, got, c.want, "should be same")
	}
}

func ReverseRunes(s string) interface{} {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
