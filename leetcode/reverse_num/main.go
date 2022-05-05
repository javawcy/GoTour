package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	v := -1212121212121212121
	fmt.Println(reverse(v))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	xStr := strconv.Itoa(x)
	newStr := []byte{}
	if x < 0 {
		newStr = append(newStr, 45)
	}
	for i := len(xStr) - 1; i >= 0; i-- {
		if i == len(xStr)-1 && xStr[i] != 48 {
			newStr = append(newStr, xStr[i])
		} else {
			if xStr[i] != 45 {
				newStr = append(newStr, xStr[i])
			}
		}
	}
	str := string(newStr)
	v, _ := strconv.Atoi(str)
	if v > int(math.Pow(2, 31))-1 || v < -int(math.Pow(2, 31)) {
		return 0
	}
	return v
}
