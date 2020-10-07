package _3_dummy_head

import (
	"bytes"
	"data_structrue/utils"
	"fmt"
)

type Node struct {
	E    interface{}
	Next *Node
}

type LinkList struct {
	DummyHead *Node
	Size      int
}

func NewLinkList() *LinkList {
	return &LinkList{DummyHead: &Node{}}
}

func (l *LinkList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkList) GetSize() int {
	return l.Size
}

func (l *LinkList) Add(index int, e interface{}) {
	if index < 0 || index > l.Size {
		panic("error index")
	}
	pre := l.DummyHead
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	n := &Node{E: e, Next: pre.Next}
	pre.Next = n
	l.Size++
}

func (l *LinkList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LinkList) AddLast(e interface{}) {
	l.Add(l.Size, e)
}

func (l *LinkList) Get(index int) interface{} {
	if index < 0 || index >= l.Size {
		panic("error index")
	}
	cur := l.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.E
}

func (l *LinkList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkList) GetLast() interface{} {
	return l.Get(l.Size - 1)
}
func (l *LinkList) Set(index int, e interface{}) {
	if index < 0 || index >= l.Size {
		panic("error index")
	}
	cur := l.DummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.E = e
}

func (l *LinkList) Contains(e interface{}) bool {
	cur := l.DummyHead.Next
	for cur != nil { //如果一个节点的下一个节点为null 说明他就是最后一个
		if utils.Compare(cur.E, e) == 0 {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (l *LinkList) Delete(index int) interface{} {
	pre := l.DummyHead
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	//找到需要删除的节点
	deleteNode := pre.Next
	pre.Next = deleteNode.Next
	deleteNode.Next = nil
	l.Size--
	return deleteNode
}

func (l *LinkList) String() string {
	buffer := bytes.Buffer{}
	cur := l.DummyHead.Next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.E) + "->")
		cur = cur.Next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
