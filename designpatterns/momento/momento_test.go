package momento

import (
	"testing"
)

func TestMomento(t *testing.T) {
	editor := NewEditor()
	history := NewHistory(3)

	editor.write("Hello")
	state := editor.createState()
	history.push(state)

	editor.write("World")
	lastState := history.pop()
	editor.restore(lastState)

	if editor.content != "Hello" {
		t.Errorf("momento pattern failed")
	}
}
