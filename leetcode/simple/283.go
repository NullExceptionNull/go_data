package simple

func moveZeroes(nums []int) {
	j := 0
	for _, value := range nums {
		if value != 0 {
			nums[j] = value
			j++
		}
	}
	for j := j; j < len(nums); j++ {
		nums[j] = 0
	}
}

func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}
