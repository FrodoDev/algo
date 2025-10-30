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

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
	// count int
	// cur   *ListNode
}

func (l *List) String() string {
	s := ""
	cur := l.Head
	for cur != nil {
		s += fmt.Sprintf("%d->", cur.Val)
		cur = cur.Next
	}
	s += "nil"
	return s
}

// 判断两个链表是否相等, 依次判断元素值是否相等, 所有元素都相等意味着两个链表相等
func (l *List) isLinkedListEqual(l2 *List) bool {
	if l == nil && l2 == nil {
		return true
	}
	if l == nil && l2 != nil {
		return false
	}
	if l != nil && l2 == nil {
		return false
	}

	cur1, cur2 := l.Head, l2.Head
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
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
func createLinkedListFromArray(arr []int) *List {
	var Head, cur *ListNode
	for i := range arr {
		if Head == nil {
			Head = &ListNode{Val: arr[i]}
			cur = Head

		} else {
			cur.Next = &ListNode{Val: arr[i]}
			cur = cur.Next
		}
	}

	l := &List{Head: Head}
	return l
}

// insertNodeHead 1. 在链表表头插入结点：空链表，只有一个节点的链表，有两个节点的链表
func (l *List) insertNodeHead(Val int) {
	if l.Head == nil {
		l.Head = &ListNode{Val: Val}
		return
	}

	newNode := &ListNode{Val: Val}
	newNode.Next = l.Head
	l.Head = newNode
}

// insertNodeTail 2. 在链表表尾插入结点：空链表，只有一个节点的链表，有两个节点的链表
func (l *List) insertNodeTail(Val int) {
	if l.Head == nil {
		l.Head = &ListNode{Val: Val}
		return
	}

	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: Val}
}

// delNodeHead 3. 在链表中删除表头节点：空链表，只有一个节点的链表，有两个节点的链表
func (l *List) delNodeHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

// delNodeTail 4. 在链表中删除表尾节点：空链表，只有一个节点的链表，有两个节点的链表
func (l *List) delNodeTail() {
	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		return
	}

	cur := l.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}

	cur.Next = nil
}

// reverse 5. 单链表反转
func (l *List) reverse() {
	var prev *ListNode
	cur := l.Head

	for cur != nil {
		Next := cur.Next
		cur.Next = prev
		prev = cur
		cur = Next
	}
	l.Head = prev
}

// hasCycle 6. 链表中环的检测
// 快慢指针,快指针每次走两步,慢指针每次走一步
func (l *List) hasCycle() bool {
	if l == nil {
		return false
	}

	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// getCycleEntrance 获取链表中环的入口结点
func (l *List) getCycleEntrance() *ListNode {
	if l == nil {
		return nil
	}

	slow, fast := l.Head, l.Head
	var node *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			node = slow
			break
		}
	}

	if node == nil {
		return nil
	}

	cur := l.Head
	for ; cur != nil; cur, slow = cur.Next, slow.Next {
		if cur == slow {
			break
		}
	}
	return cur
}

// mergeTwoSortedList 7. 有序链表合并
func mergeTwoSortedList(l1, l2 *List) *List {
	if l1 == nil || l1.Head == nil {
		return l2
	}
	if l2 == nil || l2.Head == nil {
		return l1
	}

	cur1, cur2 := l1.Head, l2.Head
	l := new(List)
	var cur *ListNode
	var min int
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			min = cur2.Val
			cur2 = cur2.Next
		} else if cur2 == nil {
			min = cur1.Val
			cur1 = cur1.Next
		} else if cur1.Val <= cur2.Val {
			min = cur1.Val
			cur1 = cur1.Next
		} else {
			min = cur2.Val
			cur2 = cur2.Next
		}

		if cur == nil {
			cur = &ListNode{Val: min}
			l.Head = cur
		} else {
			cur.Next = &ListNode{Val: min}
			cur = cur.Next
		}
	}
	return l
}

// rmListNthFromEnd 8. 删除链表倒数第 n 个结点, 超过范围的不处理
// 快慢指针,使两个指针相隔n,那么快指针到尾结点的时候,慢指针正好指向要删除节点的前一个结点
// 注意: 1. 使用辅助指针, 用一个新指针指向放在头结点前面, 这是为了方便删除头结点 2. 如果删除的正好是头结点, 头结点需要指向下一个位置
func (l *List) rmListNthFromEnd(n int) {
	if l == nil || l.Head == nil {
		return
	}

	fast := new(ListNode)
	fast.Next = l.Head
	slow := fast
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	rmNode := slow.Next
	slow.Next = slow.Next.Next
	if l.Head == rmNode { // caution! 如果删除的正好头节点, 要把头结点移走
		l.Head = slow.Next
	}
}

// getMiddleNode 9. 求链表的中间结点 (长度为偶数时,返回第一个中间结点)
// 第一种办法: 在维护链表的时候, 把长度维护上, 这样就比较方便了, 长度为 len, len为奇数时, 中间结点为len/2; len为偶数时, 中间节点为(len-1)/2
// 第二种办法: 快慢指针, 快指针每次走两步, 慢指针每次走一步, 快指针走到尾节点的时候, 慢指针正好指向中间节点
func (l *List) getMiddleNode() *ListNode {
	if l == nil {
		return nil
	}

	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// isPalindrome 10. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/
// 快慢指针,找到中间结点,中间结点到最后结点翻转,跟头结点到中间结点前一个节点部分比较
// 注意: 偶数长度的中间结点只能是第二个
// 用数字内容替代字符, 空字符串和只有一个字符的字符串被认为是回文字符串
func (l *List) isPalindrome() bool {
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return true
	}

	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// right := &list{Head: slow}
	// right.reverse()
	// cur1, cur2 := l.Head, right.Head

	var right *ListNode
	for slow != nil {
		Next := slow.Next
		slow.Next = right
		right = slow
		slow = Next
	}
	cur1, cur2 := l.Head, right

	for cur2 != nil {
		if cur2.Val != cur1.Val {
			return false
		}
		cur2 = cur2.Next
		cur1 = cur1.Next
	}
	return true
}
