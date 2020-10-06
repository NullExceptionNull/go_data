package loopqueue

import (
	"fmt"
	"testing"
)

func TestLoopQueue_DeQueue(t *testing.T) {
	queue := NewLoopQueue(10)
	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
		fmt.Println(queue)
		if i%3 == 2 {
			queue.DeQueue()
			fmt.Println(queue)
		}
	}
}
