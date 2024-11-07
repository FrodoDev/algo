package linkedlist

import "testing"

func TestIsLinkedListEqual(t *testing.T) {
	e1 := isLinkedListEqual(nil, nil)
	if !e1 {
		t.Errorf("case1: two nil linkedlist not equal?")
	}

	l1 := &listNode{}
	l2 := l1
	e2 := isLinkedListEqual(l1, l2)
	if !e2 {
		t.Errorf("case2: two same linkedlist not equal?")
	}

	l1 = &listNode{val: 1}
	l2 = l1
	e3 := isLinkedListEqual(l1, l2)
	if !e3 {
		t.Errorf("case3: two same linkedlist not equal?")
	}

	l1 = &listNode{val: 1}
	l2 = &listNode{val: 1}
	l2.next = &listNode{val: 2}
	e4 := isLinkedListEqual(l1, l2)
	if e4 {
		t.Errorf("case4: two not same linkedlist equal?")
	}

	l1 = &listNode{val: 1}
	l1.next = &listNode{val: 2}
	l2 = &listNode{val: 1}
	e5 := isLinkedListEqual(l1, l2)
	if e5 {
		t.Errorf("case5: two not same linkedlist equal?")
	}

	l1 = &listNode{val: 5}
	l1.next = &listNode{val: 2}
	l2 = &listNode{val: 2}
	l2.next = &listNode{val: 5}
	e6 := isLinkedListEqual(l1, l2)
	if e6 {
		t.Errorf("case6: two not same linkedlist equal?")
	}
}

// go test -v github/aCodeNPC/algo/linkedlist  包级测试
// go test -v -run '^Test' github/aCodeNPC/algo/linkedlist 指定函数名以Test开头的
// go test -v -run 'FromArray' github/aCodeNPC/algo/linkedlist 函数名包含字符串
// go test -v -run '^Test' github/aCodeNPC/algo/linkedlist -coverprofile=coverage.out
// go tool cover -func=coverage.out 本地查看覆盖率
// go tool cover -html=coverage.out 打开网页显示详细的覆盖情况
func TestCreateLinkedListFromArray(t *testing.T) {
	l := createLinkedListFromArray(nil)
	if l != nil {
		t.Errorf("case1: nil array create not nil linkedlist:%v", l)
	}

	l1 := createLinkedListFromArray([]int{})
	if l1 != nil {
		t.Errorf("case2: empty array create not nil linkedlist:%v", l)
	}

	l2 := createLinkedListFromArray([]int{1})
	l2c := &listNode{val: 1}
	if !isLinkedListEqual(l2, l2c) {
		t.Errorf("case3: one element array create linkedlist:%v %v", l2, l2c)
	}

	l3 := createLinkedListFromArray([]int{1, 2})
	l3c := &listNode{val: 1}
	l3c.next = &listNode{val: 2}
	if !isLinkedListEqual(l3, l3c) {
		t.Errorf("case4: two element array create linkedlist:%v %v", l3, l3c)
	}
}

func TestInsertNodeHead(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		v      int
		expect []int
	}{
		{name: "case1", arr: nil, v: 1, expect: []int{1}},
		{name: "case2", arr: []int{1}, v: 2, expect: []int{2, 1}},
		{name: "case3", arr: []int{1, 2}, v: 3, expect: []int{3, 1, 2}},
	}

	for _, c := range cases {
		l := createLinkedListFromArray(c.arr)
		ln := insertNodeHead(l, c.v)
		lc := createLinkedListFromArray(c.expect)
		if !isLinkedListEqual(ln, lc) {
			t.Errorf("%s: input:%v expect:%v ln:%v lc:%v", c.name, c.arr, c.expect, ln, lc)
		}
	}
}
