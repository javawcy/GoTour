package hashtable

import (
	"strings"
)

//HashTableError hashtable error
type HashTableError struct {
	msg string
}

func (err *HashTableError) Error() string {
	return err.msg
}

//FindFirstNoRepeatSubStr 查找第一个不重复子串
func FindFirstNoRepeatSubStr(str string) (string, error) {
	var err error
	var subStr string
	if len(str) == 0 {
		err = &HashTableError{"Str can't be empty or nil"}
	} else if len(str) == 1 {
		subStr = str
	} else {
		m := make(map[string]int)
		arr := strings.Split(str, "")
		for i := 0; i < len(arr)-1; i++ {
			if _, ok := m[arr[i]]; ok {
				delete(m, arr[i])
			} else {
				m[arr[i]] = i
			}
		}
		compare := len(str) - 1
		for k, v := range m {
			if v <= compare {
				subStr = k
				compare = v
			}
		}
	}
	return subStr, err
}
