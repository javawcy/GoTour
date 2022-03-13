package state

import "fmt"

/*
* 状态模式(比较抽象)
* 本质是对多态的使用
* Context(上下文) - State(状态) - StateImpls(状态的实现)
* context只负责发起request和切换状态，具体每个状态下的功能由stateimpls实现
* 例子：photoshop选择一个图标，鼠标样式会随之变化
 */

//Canvas context - 状态持有者
type Canvas struct {
	current Tool
}

func (c *Canvas) setCurrentTool(t Tool) {
	c.current = t
}

//context 的request方法，交由状态不同的实现
func (c *Canvas) mouseDown() {
	c.current.mouseDown()
}
func (c *Canvas) mouseUp() {
	c.current.mouseUp()
}

//Tool state - 状态
type Tool interface {
	mouseDown()
	mouseUp()
}

//SelectionTool 状态实现1
type SelectionTool struct{}

func (s *SelectionTool) mouseDown() {
	fmt.Println("selection mouse down")
}
func (s *SelectionTool) mouseUp() {
	fmt.Println("selection mouse up")
}

//DrawTool 状态实现2
type DrawTool struct{}

func (d *DrawTool) mouseDown() {
	fmt.Println("draw mouse down")
}
func (d *DrawTool) mouseUp() {
	fmt.Println("draw mouse up")
}
