package hashtable

import (
	"datastructurecase/list"
)

type entry[K Key, V any] struct {
	key   K
	value V
}

//CustomeHashMap custome map impl
type CustomeHashMap[K Key, V any] struct {
	buckets []list.LinkedList[entry[K, V]]
}

//NewCustomeHashMap init method
func NewCustomeHashMap[K Key, V any](cap int) *CustomeHashMap[K, V] {
	return &CustomeHashMap[K, V]{
		buckets: make([]list.LinkedList[entry[K, V]], cap),
	}
}

func (h *CustomeHashMap[K, V]) Put(k K, v V) {

}
