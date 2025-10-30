package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParenMatching(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"emptyString", "", true},
		{"one1-1", "()", true},
		{"one1-2", ")(", false},
		{"one1-3", "{[]", false},
		{"one1-4", "[])", false},
		{"one1-5", "[](", false},
		{"one2", "[]", true},
		{"one3", "{}", true},
		{"tow1-1", "()[]", true},
		{"tow1-2", "[()]", true},
		{"tow1-3", "([])", true},
		{"tow1-4", "[(])", false},
		{"tow2-1", "(){}", true},
		{"tow3", "[]{}", true},
		{"three3-1", "[](){}", true},
		{"three3-2", "[()]{}", true},
		{"three3-3", "{[]}()", true},
		{"three3-4", "{[()]}", true},
		{"three3-5", "{[()]}()", true},
		{"three3-6", "{([)]}()", false},
		{"three3-7", "{[()]})(", false},
		{"three3-8", "{a([b)中]国}(人)", false},
		{"three3-9", "{a([b])中国}(人)", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			get := ParenMatching(tc.input)
			assert.Equal(t, tc.want, get)
		})
	}
}
