package hashtable

import (
	"reflect"
	"testing"
)

func TestNewHashMap(t *testing.T) {
	tests := []struct {
		name string
		want *HashMap[string, int]
	}{
		{"hashmap1", &HashMap[string, int]{
			m: map[string]int{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashMap[string, int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMapIDL(t *testing.T) {
	h := NewHashMap[string, int]()
	h.Put("java", 1)
	h.Put("Golang", 2)
	if !h.ContainsKey("java") || h.ContainsKey("js") {
		t.Errorf("Contains method error")
	}
	if !h.ContainsValue(1) || h.ContainsValue(3) {
		t.Errorf("Contains value error")
	}
	v := h.Get("java")
	if v != 1 {
		t.Errorf("Get method error")
	}
}
