package arrayqueue

import (
	"bytes"
	"data_structrue/03-stacks-and-queues/queue"
	"fmt"
)

type ArrayQueue struct {
	array *queue.Array
}

func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.IsEmpty()
}

func (a *ArrayQueue) EnQueue(i interface{}) {
	a.array.AddLast(i)
}

func (a *ArrayQueue) DeQueue() interface{} {
	return a.array.RemoveFirst()
}

func (a *ArrayQueue) GetFront() interface{} {
	return a.array.GetFirst()
}

func (aq *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [")
	for i := 0; i < aq.array.GetSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(aq.array.Get(i)))
		if i != (aq.array.GetSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}

func New(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: queue.New(capacity),
	}
}
