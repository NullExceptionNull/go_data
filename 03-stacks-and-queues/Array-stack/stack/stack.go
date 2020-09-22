package stack

import (
	"bytes"
	"data_structrue/03-stacks-and-queues/Array-stack/array"
	"fmt"
)

type ArrayStack struct {
	array *array.Array
}

func New(capacity int) *ArrayStack {
	return &ArrayStack{
		array: array.New(capacity),
	}
}

func (a *ArrayStack) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayStack) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayStack) GetCapacity() int {
	return a.array.GetCapacity()
}

func (a *ArrayStack) Push(i interface{}) {
	a.array.AddLast(i)
}

func (a *ArrayStack) Pop() interface{} {
	return a.array.RemoveLast()
}

func (a *ArrayStack) Peek() interface{} {
	return a.array.GetLast()
}

func (a *ArrayStack) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Stack: ")
	buffer.WriteString("[")
	for i := 0; i < a.array.GetSize(); i++ {
		buffer.WriteString(fmt.Sprint(a.array.Get(i)))
		if i != a.array.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")

	return buffer.String()
}
