package _1

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		panic("arr size error")
	}
	//将数组的第一个节点设置为头结点
	head := &ListNode{Val: arr[0]}
	pre := head
	for i := 1; i < len(arr); i++ {
		pre.Next = &ListNode{Val: arr[i]}
		pre = pre.Next
	}
	return head
}

func (ln *ListNode) String() string {
	var buffer bytes.Buffer
	for ln != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", ln.Val))
		ln = ln.Next
	}
	buffer.WriteString("NULL")
	return buffer.String()
}
