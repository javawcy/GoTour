package hashtable

import (
	"strings"
)

type hashTableError struct {
	msg string
}

func (err *hashTableError) Error() string {
	return err.msg
}

//FindFirstNoRepeatSubStr 查找第一个不重复子串
func FindFirstNoRepeatSubStr(str string) (string, error) {
	var err error
	subStr := ""
	if len(str) == 0 {
		err = &hashTableError{"Str can't be empty or nil"}
	} else if len(str) == 1 {
		subStr = str
	} else {
		m := make(map[string]int)
		arr := strings.Split(str, "")
		for i := 0; i < len(arr); i++ {
			if _, ok := m[arr[i]]; ok {
				m[arr[i]] = m[arr[i]] + 1
			} else {
				m[arr[i]] = 1
			}
		}
		for i := 0; i < len(arr); i++ {
			if m[arr[i]] == 1 {
				subStr = arr[i]
				break
			}
		}
	}
	return subStr, err
}
