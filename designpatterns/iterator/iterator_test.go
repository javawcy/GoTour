package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	h := NewBrowseHistory[string](10)
	h.push("a")
	h.push("b")
	h.push("c")

	i := h.iterator()
	for i.hasNext() {
		fmt.Println(i.next())
	}
}
