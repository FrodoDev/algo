package practice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLink(t *testing.T) {
	cases := []testCase{
		{"nillink", nil, nil},
		{"emptylink", []int{}, []int{}},
		{"sigleElemlink", []int{1}, []int{1}},
		{"doubleElemlink", []int{1, 2}, []int{2, 1}},
		{"multiElemlink", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"repeatedElemlink", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"repeatedmultiElemlink", []int{1, 1, 2, 2}, []int{2, 2, 1, 1}},
		{"repeatedmultiElemlink1", []int{1, 2, 2, 1}, []int{1, 2, 2, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := newLink(tc.input)
			get := ReverseLink(input)
			actual := linkToSlice(get)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func newLink(arr []int) *Link {
	if arr == nil {
		return nil
	}

	flag := &Node{}
	cur := flag
	for _, elem := range arr {
		cur.next = &Node{val: elem}
		cur = cur.next
	}
	return &Link{head: flag.next}
}

func linkToSlice(l *Link) []int {
	if l == nil {
		return nil
	}

	cur := l.head
	arr := []int{}
	for cur != nil {
		arr = append(arr, cur.val)
		cur = cur.next
	}
	return arr
}

type testCase struct {
	name     string
	input    []int
	expected []int
}
