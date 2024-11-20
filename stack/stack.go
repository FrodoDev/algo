// stack 栈相关操作
// 1. 栈的基本操作: 分别用数组和链表实现创建栈, 入栈, 出栈, 栈元素数量统计
// 2. 表达式求值
// 3. 括号匹配
// 4. 实现浏览器的前进, 后退功能

package stack

import (
	"errors"
	"fmt"
	"github/aCodeNPC/algo/linkedlist"
	"regexp"
	"strconv"
	"strings"
)

// Sequential stack
// todo 实现完整的惰性操作: push,pop操作记个数,超过一定次数后,整理一下数组. 比如说, count等于2, 但栈中元素是[1, 2,    3, 4, 5], 这时应该将栈洗成[1, 2], 后面的元素是应该删除的
type SequentialStack[T myType] struct {
	stack []T
	count int
	cap   int
}

type myType interface {
	int | string
}

func NewSequentialStack[T myType](cap int) *SequentialStack[T] {
	if cap <= 0 {
		return nil
	}

	return &SequentialStack[T]{
		stack: make([]T, cap),
		count: 0,
		cap:   cap,
	}
}

// push return false means stack full
func (s *SequentialStack[T]) Push(item T) bool {
	if s.count >= s.cap {
		return false
	}

	s.stack[s.count] = item
	s.count++
	return true
}

// pop return false means stack empty
func (s *SequentialStack[T]) Pop() (T, bool) {
	var zeroVal T
	if s.count <= 0 {
		return zeroVal, false
	}

	s.count--
	v := s.stack[s.count]
	return v, true
}

// Peak 获取栈顶元素
// todo 获取栈顶运算符时,可以先peak,而不是pop,这样避免优先级高的运算符出栈再入栈
func (s *SequentialStack[T]) Peak() (T, bool) {
	var zeroVal T
	return zeroVal, false
}

func (s *SequentialStack[T]) Equal(a []T) bool {
	if s == nil && a == nil {
		return true
	}

	if s == nil {
		return false
	}

	if a == nil {
		return false
	}

	if s.count != len(a) {
		return false
	}

	for i := 0; i < s.count; i++ {
		if s.stack[i] != a[i] {
			return false
		}
	}
	return true
}

func (s *SequentialStack[T]) Count() int {
	return s.count
}

func (s *SequentialStack[T]) String() string {
	str := ""
	for i := 0; i < s.count; i++ {
		str += fmt.Sprintf("%v ", s.stack[i])
	}
	return str
}

// Chain stack
type ChainStack struct {
	count int
	stack *linkedlist.List
	tail  *linkedlist.ListNode
}

func NewChainStack() *ChainStack {
	s := new(ChainStack)

	s.stack = &linkedlist.List{}
	return s
}

func (s *ChainStack) Push(a int) {
	s.count++
	if s.stack.Head == nil {
		s.stack.Head = &linkedlist.ListNode{Val: a}
		s.tail = s.stack.Head
		return
	}

	s.tail.Next = &linkedlist.ListNode{Val: a}
	s.tail = s.tail.Next
}

func (s *ChainStack) Pop() (int, bool) {
	if s.count <= 0 {
		return 0, false
	}

	s.count--
	prev := new(linkedlist.ListNode)
	prev.Next = s.stack.Head

	for prev.Next != s.tail {
		prev = prev.Next
	}

	v := prev.Next.Val
	prev.Next = prev.Next.Next
	s.tail = prev

	// error log: 此分支写漏了, 弹空后置空
	if s.count == 0 {
		s.stack.Head = prev
	}
	return v, true
}

func (s *ChainStack) Count() int {
	return s.count
}

func (s *ChainStack) String() string {
	return s.stack.String()
}

// 2. 表达式求值 expression evaluation
// 用两个栈分别保存数字和操作符
// 在新的操作符优先级不高于上一个操作符时,弹出数字栈的两个元素用栈中最后的操作符进行运算,运算的结果再入栈

// isOp is arithmetic operator
var errInvalidExp = errors.New("expression invalid")
var errDivisionByZero = errors.New("division by zero")

func isOp(c string) bool {
	return c == "+" || c == "-" || c == "*" || c == "/" || c == "%"
}

func isDigtal(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// getOpWeight +,- weight is 1, other is 2
func getOpWeight(op string) int {
	if op == "+" || op == "-" {
		return 1
	}
	return 2
}

// opCmp compare op1 and op2 priority
func opCmp(op1, op2 string) int {
	op1W := getOpWeight(op1)
	op2W := getOpWeight(op2)
	if op1W > op2W {
		return 1
	} else if op1W < op2W {
		return -1
	}
	return 0
}

func calc(n1, n2 int, op string) (int, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errDivisionByZero
		}
		return n1 / n2, nil
	case "%":
		if n2 == 0 {
			return 0, errDivisionByZero
		}
		return n1 % n2, nil
	default:
		return 0, errInvalidExp
	}
}

func SplitExpression(expression string) []string {
	// 定义正则表达式，匹配数字和运算符
	re := regexp.MustCompile(`(\d+|[+\-*/%()])`)

	// 查找所有匹配的部分
	matches := re.FindAllString(expression, -1)

	return matches
}

// todo: 1. 0作为被除数校验(完成) 2. 用peak取代pop前一个操作符
// todo 3. 栈pop后,要修改数组元素,而不能仅仅修改count 4. 实现带括号的表达式求值
func ExpEvaluationSeq(exp string) (int, error) {
	exp = strings.ReplaceAll(exp, " ", "")
	if exp == "" {
		return 0, nil
	}

	sNum, sOp := NewSequentialStack[int](100), NewSequentialStack[string](100)
	tokens := SplitExpression(exp)
	for _, op := range tokens {
		if !(isOp(op) || isDigtal(op)) {
			return 0, errInvalidExp
		}
		fmt.Println("get op", op, isOp(op), "numStack:", sNum, "opStack:", sOp)
		if isOp(op) {
			for { // error log, 当前运算符应该跟之前已入栈的运算符循环比较优先级, 直到遇到优先级更低的或者弹空, 这是为了实现运算法则的从左到右
				preOp, ok := sOp.Pop()
				fmt.Println("pop preOp", preOp, ok, op, opCmp(preOp, op))
				if !ok {
					break
				}

				if opCmp(preOp, op) >= 0 {
					n2, ok2 := sNum.Pop()
					n1, ok1 := sNum.Pop()
					if !ok1 || !ok2 {
						return 0, errInvalidExp
					}
					n, err := calc(n1, n2, preOp)
					fmt.Println("priority calc", n1, preOp, n2, "=", n)
					if err != nil {
						return 0, err
					}
					sNum.Push(n)
				} else {
					sOp.Push(preOp)
					break
				}
			}
			sOp.Push(op)
		} else {
			n, _ := strconv.Atoi(op)
			sNum.Push(n)
		}
	}

	fmt.Println("remind numStack:", sNum, "opStack:", sOp)
	for sOp.count > 0 {
		op, ok := sOp.Pop()
		fmt.Println("pop operator", op, ok)
		if !ok {
			return 0, errInvalidExp
		}
		n2, ok2 := sNum.Pop()
		n1, ok1 := sNum.Pop()
		if !ok1 || !ok2 {
			return 0, errInvalidExp
		}
		n, err := calc(n1, n2, op)
		fmt.Println("calc", n1, op, n2, "=", n)
		if err != nil {
			return 0, err
		}
		sNum.Push(n)
	}

	if sNum.count != 1 {
		return 0, errInvalidExp
	}

	v, _ := sNum.Pop()
	return v, nil
}

// 3. 括号匹配 parenthesis matching
// func
// 4. 实现浏览器的前进, 后退功能
