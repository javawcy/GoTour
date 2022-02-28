package stack

import "strings"

//Reverse reverse string value impl by stack
func Reverse(str string) string {
	s := NewArrayStack[string](len(str))
	for _, v := range strings.Split(str, "") {
		s.Push(v)
	}
	var newStr string
	for s.Size() > 0 {
		newStr += s.Pop()
	}
	return newStr
}
