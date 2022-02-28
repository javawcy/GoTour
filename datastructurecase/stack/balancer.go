package stack

import (
	"reflect"
	"strings"
)

var leftBrackets []string = []string{"(", "<", "[", "{"}
var rightBrackets []string = []string{")", ">", "]", "}"}

func isXXBracket(s string, bs []string) bool {
	var b bool
	for _, v := range bs {
		if reflect.DeepEqual(v, s) {
			b = true
		}
	}
	return b
}

func bracketsMatch(left string, right string) bool {
	var result bool
	if (reflect.DeepEqual(left, "(") && reflect.DeepEqual(right, ")")) ||
		(reflect.DeepEqual(left, "[") && reflect.DeepEqual(right, "]")) ||
		(reflect.DeepEqual(left, "{") && reflect.DeepEqual(right, "}")) ||
		(reflect.DeepEqual(left, "<") && reflect.DeepEqual(right, ">")) {
		result = true
	}
	return result
}

//IsBalance check string contains brackets close well
func IsBalance(str string) bool {
	arr := strings.Split(str, "")
	s := NewArrayStack[string](len(str))
	for _, v := range arr {
		if isXXBracket(v, leftBrackets) {
			s.Push(v)
		}
		if isXXBracket(v, rightBrackets) && bracketsMatch(s.Peek(), v) {
			s.Pop()
		}
	}
	return s.Size() == 0
}
