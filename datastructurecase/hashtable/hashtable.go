package hashtable

import "reflect"

//HashMapError hashtable custome error
type HashMapError struct {
	msg string
}

func (err *HashMapError) Error() string {
	return err.msg
}

//Map IDL
type Map[K, V any] interface {
	Put(k K, v V)
	Get(k K) V
	ContainsKey(k K) bool
	ContainsValue(v V) bool
}

//golang map key need compareble types
type key interface {
	bool | int | int32 | int64 | float32 | float64 | string | uint | uint32 | uint64 | complex64 | complex128
}

//NewHashMap new method
func NewHashMap[K key, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		m: make(map[K]V),
	}
}

//HashMap  Map Impl
type HashMap[K key, V any] struct {
	m map[K]V
}

func (h *HashMap[K, V]) Put(k K, v V) {
	h.m[k] = v
}

func (h *HashMap[K, V]) ContainsKey(k K) bool {
	_, ok := h.m[k]
	return ok
}

func (h *HashMap[K, V]) Get(k K) V {
	return h.m[k]
}

func (h *HashMap[K, V]) ContainsValue(v V) bool {
	b := false
	for _, value := range h.m {
		if reflect.DeepEqual(value, v) {
			b = true
			break
		}
	}
	return b
}
