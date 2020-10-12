package array

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	var a = []int{0, 1, 2, 2, 3, 0, 4, 2}
	removeElement(a, 2)
	fmt.Println(a)
}
