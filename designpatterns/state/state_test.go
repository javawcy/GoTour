package state

import "testing"

func TestState(t *testing.T) {
	c := &Canvas{}
	c.setCurrentTool(&SelectionTool{})
	c.mouseDown()
	c.mouseUp()
	c.setCurrentTool(&DrawTool{})
	c.mouseDown()
	c.mouseUp()
}
