package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		if value, ok := m[v]; ok {
			return []int{value, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
