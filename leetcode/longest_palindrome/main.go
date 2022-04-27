package main

import (
	"fmt"
)

func main() {
	s := "aacabdkacaa"
	// fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome1(s))
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

	if len(s) == 0 {
		return ""
	}
	//动态规划备忘录
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	//动态方程
	// var i, j int
	// dp[i][j] = (j - i + 1 > 2) && dp[i + 1][j - 1]

	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}

	left, right := 0, 0
	max := 0

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[i] == s[j] && ((j-1)-(i+1) < 1 || dp[i+1][j-1])
			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					left = i
					right = j
				}
			}
		}
	}
	return s[left : right+1]
}
