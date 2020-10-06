package array_stack

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{} //获取栈尾的元素
}
