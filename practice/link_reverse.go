// 链表反转

package practice

func ReverseLink(link *Link) *Link {
	if link == nil {
		return nil
	}

	head := link.head
	var prev *Node
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}

	return &Link{head: prev}
}

type Node struct {
	val  int
	next *Node
}

type Link struct {
	head *Node
}
