package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrReverse(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string
	}{
		{"case1", "", ""},
		{"case2", "abc", "cba"},
		{"case3", " 12ab ", " ba21 "},
		{"case4", "~1a!", "!a1~"},
		{"case5", "中国人", "人国中"},
		{"case4", "hello 中国人", "人国中 olleh"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := ReverseStr(c.src)
			assert.Equal(t, c.want, get)
		})
	}
}
