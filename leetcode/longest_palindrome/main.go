package main

import (
	"fmt"
)

func main() {
	s := "ccc"
	fmt.Println(longestPalindrome(s))
}

//中心扩散
func longestPalindrome(s string) string {
	arr := []byte(s)
	max := 0
	res := []byte{}
	for i := 0; i < len(s); i++ {
		if i-1 < 0 && i+1 >= len(s) {
			return string(arr)
		}
		if i+1 < len(s) && arr[i] == arr[i+1] {
			c := check(i, i+1, arr)
			if len(c) > max {
				max = len(c)
				res = c
			}
		}
		if i-1 >= 0 && i+1 < len(s) && arr[i-1] == arr[i+1] {
			c := check(i, i, arr)
			if len(c) > max {
				max = len(c)
				res = c
			}
		}
	}
	if max == 0 {
		res = arr[0:1]
	}
	return string(res)
}

func check(start, end int, arr []byte) []byte {
	for start >= 0 && end < len(arr) {
		if arr[start] == arr[end] {
			start--
			end++
		} else {
			break
		}
	}
	return arr[start+1 : end]
}

//动态规划
func longestPalindrome1(s string) string {

}
