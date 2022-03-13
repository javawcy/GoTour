package momento

import "datastructurecase/stack"

/*
 *  备忘录模式 - 解决回撤问题
 *  角色分为 Orignator(发起者) - Momento(备忘录) - CareTaker(监护者)
 *  发起者每次发起一个行为的同时记到备忘录，由监护者保存备忘录，发起者回撤操作时向监护者要上一个备忘录
 *  例子：编辑器回撤
 */

//Editor 编辑器
type Editor struct {
	content string
}

//NewEditor init method
func NewEditor() *Editor {
	return &Editor{}
}

//写
func (e *Editor) write(content string) {
	e.content = content
}

//创建副本状态
func (e *Editor) createState() *editorState {
	return &editorState{e.content}
}

//还原状态
func (e *Editor) restore(s *editorState) {
	e.content = s.content
}

type editorState struct {
	content string
}

//History 历史记录使用栈保存状态保证先进后出
type History struct {
	states stack.Stack[*editorState]
}

//NewHistory init method
func NewHistory(cap int) *History {
	s := stack.NewArrayStack[*editorState](cap)
	return &History{states: s}
}

func (h *History) push(s *editorState) {
	h.states.Push(s)
}

func (h *History) pop() *editorState {
	return h.states.Pop()
}
