// 实现如下链表功能
// 1. 在链表表头插入结点：空链表，只有一个节点的链表，有两个节点的链表
// 2. 在链表表尾插入结点：空链表，只有一个节点的链表，有两个节点的链表
// 3. 在链表中删除表头节点：空链表，只有一个节点的链表，有两个节点的链表
// 4. 在链表中删除表尾节点：空链表，只有一个节点的链表，有两个节点的链表
// 3. 单链表反转
// 4. 链表中环的检测
// 5. 有序链表合并
// 6. 删除链表倒数第 n 个结点
// 7. 求链表的中间结点
// 8. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/

package linkedlist

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

type list struct {
	head *listNode
	// count int
	// cur   *listNode
}

func (l *list) String() string {
	s := ""
	cur := l.head
	for cur != nil {
		if s == "" {
			s += fmt.Sprintf("%d", cur.val)
		} else {
			s += fmt.Sprintf(" %d", cur.val)
		}
		cur = cur.next
	}
	return s
}

// 判断两个链表是否相等, 依次判断元素值是否相等, 所有元素都相等意味着两个链表相等
func (l *list) isLinkedListEqual(l2 *list) bool {
	if l == nil && l2 == nil {
		return true
	}
	if l == nil && l2 != nil {
		return false
	}
	if l != nil && l2 == nil {
		return false
	}

	cur1, cur2 := l.head, l2.head
	for cur1 != nil && cur2 != nil {
		if cur1.val != cur2.val {
			return false
		}
		cur1 = cur1.next
		cur2 = cur2.next
	}

	if cur1 != nil {
		return false
	}
	if cur2 != nil {
		return false
	}
	return true
}

// createLinkedListFromArray 创建一个链表, 返回链表头结点. 链表中的节点值就是数组元素依次赋值.
func createLinkedListFromArray(arr []int) *list {
	var head, cur *listNode
	for i := range arr {
		if head == nil {
			head = &listNode{val: arr[i]}
			cur = head

		} else {
			cur.next = &listNode{val: arr[i]}
			cur = cur.next
		}
	}

	l := &list{head: head}
	return l
}

// insertNodeHead 1. 在链表表头插入结点：空链表，只有一个节点的链表，有两个节点的链表
func (l *list) insertNodeHead(val int) {
	if l.head == nil {
		l.head = &listNode{val: val}
		return
	}

	newNode := &listNode{val: val}
	newNode.next = l.head
	l.head = newNode
}

// insertNodeTail 2. 在链表表尾插入结点：空链表，只有一个节点的链表，有两个节点的链表
func (l *list) insertNodeTail(val int) {
	if l.head == nil {
		l.head = &listNode{val: val}
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &listNode{val: val}
}

// delNodeHead 3. 在链表中删除表头节点：空链表，只有一个节点的链表，有两个节点的链表
func (l *list) delNodeHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

// 4. 在链表中删除表尾节点：空链表，只有一个节点的链表，有两个节点的链表
func (l *list) delNodeTail() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	cur := l.head
	for cur.next.next != nil {
		cur = cur.next
	}

	cur.next = nil
}

// 3. 单链表反转
// 4. 链表中环的检测
// 5. 有序链表合并
// 6. 删除链表倒数第 n 个结点
// 7. 求链表的中间结点
// 8. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/
