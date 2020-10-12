package simple

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return nil
	}
	var i = 0
	var j = len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
