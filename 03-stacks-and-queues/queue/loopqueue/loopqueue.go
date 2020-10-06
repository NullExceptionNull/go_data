package loopqueue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	Front int
	Tail  int
	Data  []interface{}
	Size  int
}

func NewLoopQueue(capacity int) *LoopQueue {
	return &LoopQueue{
		Data: make([]interface{}, capacity+1),
	}
}

func (a *LoopQueue) GetCapacity() int {
	return len(a.Data) - 1
}

func (a *LoopQueue) IsEmpty() bool {
	return a.Front == a.Tail
}

func (a *LoopQueue) GetSize() int {
	return a.Size
}

func (a *LoopQueue) EnQueue(e interface{}) {
	//首先得判断队列是不是已经满了 对于循环队列 如果 (tail+1)%C == front
	if (a.Tail+1)%len(a.Data) == a.Front {
		a.resize(a.GetCapacity() * 2)
	}
	a.Data[a.Tail] = e
	a.Tail = (a.Tail + 1) % len(a.Data)
	a.Size++
}

func (a *LoopQueue) DeQueue() interface{} {
	if a.IsEmpty() {
		panic("Queue is empty")
	}
	e := a.Data[a.Front]
	a.Data[a.Front] = nil
	a.Front = (a.Front + 1) % len(a.Data) //出队的时候需要把队首向前挪一个位置
	a.Size--
	if a.Size == a.GetCapacity()/4 && a.GetCapacity()/2 != 0 {
		a.resize(a.GetCapacity() / 2)
	}
	return e
}

func (a *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity+1)
	for i := 0; i < a.Size; i++ {
		newData[i] = a.Data[(i+a.Front)%len(a.Data)]
	}
	a.Data = newData
	a.Front = 0
	a.Tail = a.Size
}

func (a *LoopQueue) GetFront() interface{} {
	if a.IsEmpty() {
		panic("Queue is empty")
	}

	return a.Data[a.Front]
}

func (lq *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", lq.Size, lq.GetCapacity()))
	buffer.WriteString("front [")
	for i := lq.Front; i != lq.Tail; i = (i + 1) % len(lq.Data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(lq.Data[i]))
		if (i+1)%len(lq.Data) != lq.Tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
