// 判断表达式中，括号是否匹配
package practice

/*
使用栈，存储左括号，遇到右括号，判断跟栈顶元素是否匹配
1. 如果栈为空，说明表达式括号不匹配
2. 如果不匹配，说明表达式括号不匹配
3. 遍历完字符串，如果栈中还有元素，说明表达式括号不匹配
*/
func ParenMatching(s string) bool {
	stack := []rune{}
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			// 入栈
			stack = append(stack, rune(s[i]))
		case ')', ']', '}':
			// 比对栈顶
			if len(stack) <= 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (s[i] == ')' && top == '(') ||
				(s[i] == ']' && top == '[') ||
				(s[i] == '}' && top == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			// 忽略
		}
	}
	// 栈空，匹配
	return len(stack) <= 0
}
