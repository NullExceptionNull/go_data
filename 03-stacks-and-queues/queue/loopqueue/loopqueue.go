package loopqueue

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
