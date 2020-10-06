package queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	EnQueue(interface{})
	DeQueue() interface{}
	GetFront() interface{}
}
