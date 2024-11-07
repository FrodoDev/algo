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

type listNode struct {
	val  int
	next *listNode
}

// 判断两个链表是否相等, 依次判断元素值是否相等, 所有元素都相等意味着两个链表相等
func isLinkedListEqual(l1, l2 *listNode) bool {
	for l1 != nil && l2 != nil {
		if l1.val != l2.val {
			return false
		}
		l1 = l1.next
		l2 = l2.next
	}

	if l1 != nil {
		return false
	}
	if l2 != nil {
		return false
	}
	return true
}

// createLinkedListFromArray 创建一个链表, 返回链表头结点. 链表中的节点值就是数组元素依次赋值.
func createLinkedListFromArray(arr []int) *listNode {
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
	return head
}

// insertNodeHead 1. 在链表表头插入结点：空链表，只有一个节点的链表，有两个节点的链表
func insertNodeHead(list *listNode, val int) *listNode {
	if list == nil {
		return &listNode{val: val}
	}

	newNode := &listNode{val: val}
	newNode.next = list
	return newNode
}

// 2. 在链表表尾插入结点：空链表，只有一个节点的链表，有两个节点的链表
// 3. 在链表中删除表头节点：空链表，只有一个节点的链表，有两个节点的链表
// 4. 在链表中删除表尾节点：空链表，只有一个节点的链表，有两个节点的链表
// 3. 单链表反转
// 4. 链表中环的检测
// 5. 有序链表合并
// 6. 删除链表倒数第 n 个结点
// 7. 求链表的中间结点
// 8. 判断一个字符串是否是回文字符串（链表）Palindrome string /ˈpæl.ɪn.droʊm/
