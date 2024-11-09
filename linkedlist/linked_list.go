// 实现如下链表功能
// 1. 在链表表头插入结点：空链表，只有一个节点的链表，有两个节点的链表
// 2. 在链表表尾插入结点：空链表，只有一个节点的链表，有两个节点的链表
// 3. 在链表中删除表头节点：空链表，只有一个节点的链表，有两个节点的链表
// 4. 在链表中删除表尾节点：空链表，只有一个节点的链表，有两个节点的链表
// 5. 单链表反转
// 6. 链表中环的检测
// 7. 有序链表合并
// 8. 删除链表倒数第 n 个结点
// 9. 求链表的中间结点
// 10. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/

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
		s += fmt.Sprintf("%d->", cur.val)
		cur = cur.next
	}
	s += "nil"
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

// reverse 3. 单链表反转
func (l *list) reverse() {
	var prev *listNode
	cur := l.head

	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

// hasCycle 4. 链表中环的检测
// 快慢指针,快指针每次走两步,慢指针每次走一步
func (l *list) hasCycle() bool {
	if l == nil {
		return false
	}

	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if fast == slow {
			return true
		}
	}

	return false
}

func (l *list) getCycleEntrance() *listNode {
	if l == nil {
		return nil
	}

	slow, fast := l.head, l.head
	var node *listNode
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if fast == slow {
			node = slow
			break
		}
	}

	if node == nil {
		return nil
	}

	cur := l.head
	for ; cur != nil; cur, slow = cur.next, slow.next {
		if cur == slow {
			break
		}
	}
	return cur
}

// mergeTwoSortedList 5. 有序链表合并
func mergeTwoSortedList(l1, l2 *list) *list {
	if l1 == nil || l1.head == nil {
		return l2
	}
	if l2 == nil || l2.head == nil {
		return l1
	}

	cur1, cur2 := l1.head, l2.head
	l := new(list)
	var cur *listNode
	var min int
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			min = cur2.val
			cur2 = cur2.next
		} else if cur2 == nil {
			min = cur1.val
			cur1 = cur1.next
		} else if cur1.val <= cur2.val {
			min = cur1.val
			cur1 = cur1.next
		} else {
			min = cur2.val
			cur2 = cur2.next
		}

		if cur == nil {
			cur = &listNode{val: min}
			l.head = cur
		} else {
			cur.next = &listNode{val: min}
			cur = cur.next
		}
	}
	return l
}

// rmListNthFromEnd 6. 删除链表倒数第 n 个结点, 超过范围的不处理
// 快慢指针,使两个指针相隔n,那么快指针到尾结点的时候,慢指针正好指向要删除节点的前一个结点
// 注意: 1. 使用辅助指针, 用一个新指针指向放在头结点前面, 这是为了方便删除头结点 2. 如果删除的正好是头结点, 头结点需要指向下一个位置
func (l *list) rmListNthFromEnd(n int) {
	if l == nil || l.head == nil {
		return
	}

	fast := new(listNode)
	fast.next = l.head
	slow := fast
	for i := 0; i < n; i++ {
		fast = fast.next
		if fast == nil {
			return
		}
	}

	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	rmNode := slow.next
	slow.next = slow.next.next
	if l.head == rmNode { // caution! 如果删除的正好头节点, 要把头结点移走
		l.head = slow.next
	}
}

// 7. 求链表的中间结点 (长度为偶数时,返回第一个中间结点)
// 第一种办法: 在维护链表的时候, 把长度维护上, 这样就比较方便了, 长度为 len, len为奇数时, 中间结点为len/2; len为偶数时, 中间节点为(len-1)/2
// 第二种办法: 快慢指针, 快指针每次走两步, 慢指针每次走一步, 快指针走到尾节点的时候, 慢指针正好指向中间节点
func (l *list) getMiddleNode() *listNode {
	if l == nil {
		return nil
	}

	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// 8. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/
// 快慢指针,找到中间结点,中间结点到最后结点翻转,跟头结点到中间结点前一个节点部分比较
// 注意: 偶数长度的中间结点只能是第二个
// 用数字内容替代字符, 空字符串和只有一个字符的字符串被认为是回文字符串
func (l *list) isPalindrome() bool {
	if l == nil || l.head == nil || l.head.next == nil {
		return true
	}

	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// right := &list{head: slow}
	// right.reverse()
	// cur1, cur2 := l.head, right.head

	var right *listNode
	for slow != nil {
		next := slow.next
		slow.next = right
		right = slow
		slow = next
	}
	cur1, cur2 := l.head, right

	for cur2 != nil {
		if cur2.val != cur1.val {
			return false
		}
		cur2 = cur2.next
		cur1 = cur1.next
	}
	return true
}
