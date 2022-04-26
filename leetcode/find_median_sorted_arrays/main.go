package main

import "fmt"

func main() {

	//前后
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	// fmt.Println(findMedianSortedArrays(nums1, nums2))
	// fmt.Println(findMedianSortedArrays(nums2, nums1))
	// //包含
	// nums1 = []int{1, 3}
	// nums2 = []int{2}
	// fmt.Println(findMedianSortedArrays(nums1, nums2))
	// // 交叉
	// nums1 = []int{1, 2, 5, 6}
	// nums2 = []int{3, 7}
	// fmt.Println(findMedianSortedArrays(nums1, nums2))
	//相等
	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newArr := []int{}
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	} else if len(nums1) == 0 {
		newArr = nums2
	} else if len(nums2) == 0 {
		newArr = nums1
	} else {
		//拼接数组
		a1 := nums1[0]
		a2 := nums1[len(nums1)-1]
		b1 := nums2[0]
		b2 := nums2[len(nums2)-1]

		if a1 == a2 && b1 == b2 && a1 == b1 {
			newArr = append(newArr, nums1...)
			newArr = append(newArr, nums2...)
		} else {
			//前后
			if b1 > a2 {
				newArr = append(newArr, nums1...)
				newArr = append(newArr, nums2...)
			}

			if a1 > b2 {
				newArr = append(newArr, nums2...)
				newArr = append(newArr, nums1...)
			}
			//组合
			if b1 >= a1 && b1 <= a2 {
				//交叉还是包含
				if b2 <= a2 {
					//包含
					newArr = cover(nums1, nums2, a1, b1, a2, b2)
				} else {
					//交叉
					newArr = x(nums1, nums2, a1, b1, a2, b2)
				}
			}

			if a1 > b1 && a1 < b2 {
				if a2 < b2 {
					//包含
					newArr = cover(nums2, nums1, b1, a1, b2, a2)
				} else {
					//交叉
					newArr = x(nums2, nums1, b1, a1, b2, a2)
				}
			}
		}
	}

	var min, max float64
	if len(newArr)%2 == 0 {
		if len(newArr) == 0 {
			return 0
		} else if len(newArr) == 2 {
			min = float64(newArr[0])
			max = float64(newArr[len(newArr)-1])
			return (min + max) / 2
		} else {
			middle_index := (len(newArr) - 1) / 2
			min = float64(newArr[middle_index])
			max = float64(newArr[middle_index+1])
			return (min + max) / 2
		}
	} else {
		if len(newArr) == 1 {
			return float64(newArr[0])
		} else {
			middle_index := (len(newArr) - 1) / 2
			return float64(newArr[middle_index])
		}
	}
}

func cover(nums1, nums2 []int, a1, b1, a2, b2 int) []int {
	a_index := 0
	b_index := 0
	newArr := []int{}
	for i := 0; nums1[i] <= b1; i++ {
		newArr = append(newArr, nums1[i])
		a_index = i
	}
	for b_index < len(nums2) {
		if nums1[a_index+1] < nums2[b_index] {
			newArr = append(newArr, nums1[a_index+1])
			a_index++
		} else {
			newArr = append(newArr, nums2[b_index])
			b_index++
		}
	}
	for i := a_index + 1; i < len(nums1); i++ {
		newArr = append(newArr, nums1[i])
	}
	return newArr
}

func x(nums1, nums2 []int, a1, b1, a2, b2 int) []int {
	a_index := 0
	b_index := 0
	newArr := []int{}
	for i := 0; nums1[i] < b1; i++ {
		newArr = append(newArr, nums1[i])
		a_index = i
	}
	for a_index+1 < len(nums1) {
		if nums1[a_index+1] < nums2[b_index] {
			newArr = append(newArr, nums1[a_index+1])
			a_index++
		} else {
			newArr = append(newArr, nums2[b_index])
			b_index++
		}
	}
	for i := b_index; i < len(nums2); i++ {
		newArr = append(newArr, nums2[i])
	}
	return newArr
}
