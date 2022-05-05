package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 2))
}

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	if numRows < 2 {
		return s
	}
	res := make([][]byte, len(s))
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, numRows)
	}
	col := 1
	cap := 0
	current := 0
	for current < len(s) {
		if cap < numRows {
			if col%2 == 0 {
				y := (numRows - 1) - cap - 1
				if y > 0 {
					res[col-1][y] = s[current]
					current++
					cap++
				} else {
					col++
					cap = 0
				}
			} else {
				res[col-1][cap] = s[current]
				current++
				cap++
			}
		} else {
			col++
			cap = 0
		}
	}

	r := []byte{}
	currentLine := 0
	currentCol := 0
	for currentLine < numRows {
		if res[currentCol][currentLine] > 0 {
			r = append(r, res[currentCol][currentLine])
		}
		currentCol++
		if currentCol == len(res) {
			currentLine++
			currentCol = 0
		}
	}
	return string(r)
}
