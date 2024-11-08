package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v github/aCodeNPC/algo/linkedlist  包级测试
// go test -v -run '^Test' github/aCodeNPC/algo/linkedlist 指定函数名以Test开头的
// go test -v -run 'FromArray' github/aCodeNPC/algo/linkedlist 函数名包含字符串
// go test -v -run '^Test' github/aCodeNPC/algo/linkedlist -coverprofile=coverage.out
// go tool cover -func=coverage.out 本地查看覆盖率
// go tool cover -html=coverage.out 打开网页显示详细的覆盖情况

func TestIsLinkedListEqual(t *testing.T) {
	l1 := new(list)
	e1 := l1.isLinkedListEqual(nil)
	if e1 {
		t.Errorf("case1: empty list equal nil list")
	}

	l1 = new(list)
	l2 := l1
	e2 := l1.isLinkedListEqual(l2)
	if !e2 {
		t.Errorf("case2: two same linkedlist not equal?")
	}

	l1 = new(list)
	l1.head = &listNode{val: 1}
	l2 = l1
	e3 := l1.isLinkedListEqual(l2)
	if !e3 {
		t.Errorf("case3: two same linkedlist not equal?")
	}

	l1 = new(list)
	l1.head = &listNode{val: 1}
	l2 = new(list)
	l2.head = &listNode{val: 1}
	l2.head.next = &listNode{val: 2}
	e4 := l1.isLinkedListEqual(l2)
	if e4 {
		t.Errorf("case4: two not same linkedlist equal?")
	}

	l1 = new(list)
	l1.head = &listNode{val: 1}
	l1.head.next = &listNode{val: 2}
	l2 = new(list)
	l2.head = &listNode{val: 1}

	e5 := l1.isLinkedListEqual(l2)
	if e5 {
		t.Errorf("case5: two not same linkedlist equal?")
	}

	l1 = createLinkedListFromArray([]int{5, 2})
	l2 = createLinkedListFromArray([]int{2, 5})
	e6 := l1.isLinkedListEqual(l2)
	if e6 {
		t.Errorf("case6: two not same linkedlist equal?")
	}

	var l *list
	e7 := l.isLinkedListEqual(nil)
	assert.Equal(t, true, e7)

	e8 := l.isLinkedListEqual(l2)
	assert.NotEqual(t, true, e8)

}

func TestCreateLinkedListFromArray(t *testing.T) {
	l := createLinkedListFromArray(nil)
	if l.head != nil {
		t.Errorf("nil arr create not nil list.head")
	}
	assert.Equal(t, nil, nil)

	l1 := createLinkedListFromArray([]int{})
	if l1.head != nil {
		t.Errorf("empty arr create not nil list.head")
	}
	// assert.Equal(t, nil, l1)

	l2 := createLinkedListFromArray([]int{1})
	l2c := new(list)
	l2c.head = &listNode{val: 1}
	if !l2.isLinkedListEqual(l2c) {
		t.Errorf("case3: one element array create linkedlist:%v %v", l2, l2c)
	}

	l3 := createLinkedListFromArray([]int{1, 2})
	l3c := new(list)
	l3c.head = &listNode{val: 1}
	l3c.head.next = &listNode{val: 2}
	if !l3.isLinkedListEqual(l3c) {
		t.Errorf("case4: two element array create linkedlist:%v %v", l3, l3c)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{"test1", nil, ""},
		{"test2", []int{}, ""},
		{"test3", []int{1}, "1"},
		{"test4", []int{1, 2}, "1 2"},
		{"test5", []int{1, 2, 3}, "1 2 3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := createLinkedListFromArray(tt.input)
			s := l.String()
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestInsertNodeHead(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		insert int
		want   []int
	}{
		{name: "case1", input: nil, insert: 1, want: []int{1}},
		{name: "case2", input: []int{}, insert: 1, want: []int{1}},
		{name: "case3", input: []int{1}, insert: 2, want: []int{2, 1}},
		{name: "case4", input: []int{1, 2}, insert: 3, want: []int{3, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := createLinkedListFromArray(tt.input)
			li.insertNodeHead(tt.insert)
			lw := createLinkedListFromArray(tt.want)
			if !li.isLinkedListEqual(lw) {
				t.Errorf("%s: input:%v expect:%v li:%v lw:%v", tt.name, tt.input, tt.want, li, lw)
			}
		})
	}
}

func TestInsertNodeTail(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		insert int
		want   []int
	}{
		{"test1", nil, 1, []int{1}},
		{"test2", []int{}, 1, []int{1}},
		{"test3", []int{1}, 2, []int{1, 2}},
		{"test4", []int{1, 2}, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := createLinkedListFromArray(tt.input)
			li.insertNodeTail(tt.insert)
			lw := createLinkedListFromArray(tt.want)
			if !li.isLinkedListEqual(lw) {
				t.Errorf("%s: input:%v expect:%v li:%v lw:%v", tt.name, tt.input, tt.want, li, lw)
			}
		})

	}
}

func TestDelNodeHead(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"test1", nil, nil},
		{"test2", []int{}, nil},
		{"test3", []int{1}, nil},
		{"test4", []int{1, 2}, []int{2}},
		{"test5", []int{1, 2, 3}, []int{2, 3}},
		{"test5", []int{1, 2, 3, 4}, []int{2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := createLinkedListFromArray(tt.input)
			li.delNodeHead()
			lw := createLinkedListFromArray(tt.want)
			if !li.isLinkedListEqual(lw) {
				t.Errorf("li:%v != lw:%v", li, lw)
			}
		})
	}
}

func TestDelNodeTail(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"test1", nil, nil},
		{"test2", []int{}, nil},
		{"test3", []int{1}, nil},
		{"test4", []int{1, 2}, []int{1}},
		{"test5", []int{1, 2, 3}, []int{1, 2}},
		{"test5", []int{1, 2, 3, 4}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := createLinkedListFromArray(tt.input)
			li.delNodeTail()
			lw := createLinkedListFromArray(tt.want)
			if !li.isLinkedListEqual(lw) {
				t.Errorf("li:%v != lw:%v", li, lw)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"test1", nil, nil},
		{"test2", []int{}, []int{}},
		{"test3", []int{1}, []int{1}},
		{"test4", []int{1, 2}, []int{2, 1}},
		{"test5", []int{1, 2, 3}, []int{3, 2, 1}},
		{"test6", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := createLinkedListFromArray(tt.input)
			lw := createLinkedListFromArray(tt.want)
			li.reverse()
			if !li.isLinkedListEqual(lw) {
				t.Errorf("input:%v want:%v inputlist reverse:%v wantlist:%v", tt.input, tt.want, li, lw)
			}
		})
	}
}
