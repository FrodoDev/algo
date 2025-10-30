package practice

func ReverseStr(s string) string {
	newstr := []rune(s)
	for i, j := 0, len(newstr)-1; i < j; i, j = i+1, j-1 {
		newstr[i], newstr[j] = newstr[j], newstr[i]
	}
	return string(newstr)
}
