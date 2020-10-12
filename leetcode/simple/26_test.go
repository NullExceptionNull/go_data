package simple

import (
	"fmt"
	"testing"
)

func Test26(t *testing.T) {
	var nums = []int{1, 2, 3, 3, 4, 4}

	duplicates := removeDuplicates(nums)

	fmt.Println(duplicates)

}
