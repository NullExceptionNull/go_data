package _2_add_element

import "fmt"

type Node struct {
	Next *Node
	E    interface{}
}

type LinkList struct {
	Head *Node
	Size int
}

func (l *LinkList) IsEmpty() bool {
	return l.Size == 0
}

func (l *LinkList) GetSize() int {
	return l.Size
}

//添加一个头结点 那被添加的节点的下一个节点就是之前的头结点
func (l *LinkList) AddFirst(e interface{}) {
	node := &Node{E: e}
	node.Next = l.Head
	l.Head = node
	l.Size++
}

func (l *LinkList) Add(index int, e interface{}) {
	if index < 0 || index > l.Size {
		panic("error index")
	}
	if index == 0 {
		l.AddFirst(e)
	} else {
		pre := l.Head
		for i := 0; i < index-1; i++ {
			pre = pre.Next
		}
		n := &Node{E: e, Next: pre.Next}
		pre.Next = n
		l.Size++
	}
}

func (l *LinkList) AddLast(e interface{}) {
	l.Add(l.Size, e)
}

func (n *Node) String() string {
	return fmt.Sprint(n.E)
}
