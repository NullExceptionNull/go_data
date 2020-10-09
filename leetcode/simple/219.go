package simple

import (
	"math"
)

//给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引i和j，使得nums [i] = nums [j]，并且 i 和 j的差的 绝对值 至多为 k。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/contains-duplicate-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1 == num2 && i > j {
				m[i] = j
			}
		}
	}
	if len(m) == 0 {
		return false
	}
	var ret []float64
	for k, v := range m {
		ret = append(ret, math.Abs(float64(k-v)))

	}
	for _, f := range ret {
		if int(f) <= k {
			return true
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {

	m := make(map[int]int)
	for i, num := range nums {
		if i2, ok := m[num]; ok && i-i2 <= k {
			return true
		}
		m[num] = i
	}
	return false
}
