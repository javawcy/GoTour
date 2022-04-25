package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("aabc"))
}

func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	right := -1
	max := 0
	for left := 0; left < len(s); left++ {
		//左边界移动一格删除前一格数据
		if left > 0 {
			delete(window, s[left-1])
		}

		//移动右边届, 条件是移动目标位置的元素不在窗口内也就是不重复
		for right+1 < len(s) && window[s[right+1]] == 0 {
			right++
			window[s[right]] = 1
		}

		//记录每次窗口遇到重复元素或到底时的长度
		l := (right - left) + 1
		if l > max {
			max = l
		}
	}
	return max
}
