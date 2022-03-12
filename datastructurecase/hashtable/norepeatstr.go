package hashtable

import (
	"strings"
)

//FindFirstNoRepeatSubStr 查找第一个不重复子串
func FindFirstNoRepeatSubStr(str string) (string, error) {
	var err error
	subStr := ""
	if len(str) == 0 {
		err = &HashMapError{"Str can't be empty or nil"}
	} else if len(str) == 1 {
		subStr = str
	} else {
		h := NewHashMap[string, int]()
		arr := strings.Split(str, "")
		for i := 0; i < len(arr); i++ {
			if h.ContainsKey(arr[i]) {
				v := h.Get(arr[i])
				h.Put(arr[i], v+1)
			} else {
				h.Put(arr[i], 1)
			}
		}
		for i := 0; i < len(arr); i++ {
			if h.Get(arr[i]) == 1 {
				subStr = arr[i]
				break
			}
		}
	}
	return subStr, err
}
