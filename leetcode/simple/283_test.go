package simple

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {

	var nums = []int{1, 0, 2}
	moveZeroes(nums)
	fmt.Println(nums)

}
