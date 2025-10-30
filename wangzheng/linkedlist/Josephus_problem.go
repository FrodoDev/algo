// Josephus Problem
// 约瑟夫斯置换是一个出现在计算机科学和数学中的问题。在计算机编程的算法中，类似问题又被称为约瑟夫环。
// 人们站在一个等待被处决的圈子里。 计数从圆圈中的指定点开始，并沿指定方向围绕圆圈进行。
// 在跳过指定数量的人之后，处刑下一个人。 对剩下的人重复该过程，从下一个人开始，朝同一方向跳过相同数量的人，直到只剩下一个人，并被释放。
// 问题即，给定人数、起点、方向和要跳过的数字，选择初始圆圈中的位置以避免被处决。

package linkedlist

import "fmt"

// JosephusProblem 处理约瑟夫问题, n为总人数, k为每次数的人, 返回值为 survivor 编号
func JosephusProblem(n, k int) int {
	circle := newCircleList(n)
	if circle == nil {
		return -1
	}

	// fmt.Printf("(%d %d): %v\n", n, k, circle)

	preN := k - 1
	remindN := n
	cur := circle.tail
	for remindN > 1 {
		for i := 0; i < preN; i++ {
			cur = cur.Next
		}

		cur.Next = cur.Next.Next
		remindN -= 1
	}

	circle.head = cur.Next
	circle.tail = cur
	return cur.Val
}

type circleList struct {
	head, tail *ListNode
}

func (l *circleList) String() string {
	s := ""
	cur := l.head
	for cur != l.tail {
		s += fmt.Sprintf("%d->", cur.Val)
		cur = cur.Next
	}
	s += fmt.Sprintf("%d->%d", cur.Val, cur.Next.Val)
	return s
}

func newCircleList(n int) *circleList {
	if n <= 0 {
		return nil
	}

	l := new(circleList)
	var cur *ListNode
	for i := 0; i < n; i++ {
		if l.head == nil {
			l.head = &ListNode{Val: i + 1} // 这一步很重要, 要把head初始化
			cur = l.head
		} else {
			cur.Next = &ListNode{Val: i + 1}
			cur = cur.Next
		}
	}
	cur.Next = l.head
	l.tail = cur
	// fmt.Println("new get", l.head, l.tail.Val)
	return l
}
