package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	if len(s) == 0 {
		return 0
	}

	arr := strings.Split(s, "")
	i := 0
	j := 0
	m := make(map[string]int, len(arr))
	for i < len(arr) && j < len(arr) {
		if _, ok := m[arr[j]]; ok {
			l := len(m)
			if l > max {
				max = l
			}
			m = make(map[string]int, len(arr))
			i++
			j = i
		} else {
			m[arr[j]] = j
			j++
			if j == len(arr) {
				l := len(m)
				if l > max {
					max = l
				}
				m = make(map[string]int, len(arr))
				i++
				j = i
			}
		}
	}
	return max
}
