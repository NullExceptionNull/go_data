package simple

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {

	var nums = []int{1, 2, 3, 1, 2, 3}

	duplicate := containsNearbyDuplicate(nums, 2)

	fmt.Println(duplicate)
}
