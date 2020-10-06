package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New(10)
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
