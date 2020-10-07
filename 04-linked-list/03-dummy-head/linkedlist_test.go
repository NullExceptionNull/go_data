package _3_dummy_head

import (
	"fmt"
	"testing"
)

func TestLinkList_Add(t *testing.T) {
	linkedList := NewLinkList()

	for i := 0; i < 5; i++ {
		linkedList.AddFirst(i)
		fmt.Println(linkedList)
	}
	linkedList.Add(2, 666)
	fmt.Println(linkedList)
	i := linkedList.Delete(2)
	fmt.Println(i.(*Node).E)
	fmt.Println(linkedList)
}
