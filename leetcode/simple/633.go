package simple

import "math"

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。

func judgeSquareSum(c int) bool {

	if c == 0 {
		return true
	}

	var i = 0
	var j = int(math.Sqrt(float64(c)))

	for i <= j {
		if i*i+j*j == c {
			return true
		} else if i*i+j*j <= c {
			i++
		} else {
			j--
		}
	}
	return false
}
