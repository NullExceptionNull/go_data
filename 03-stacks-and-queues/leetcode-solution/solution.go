package leetcode_solution

func isValid(s string) bool {

	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}

	var stack []rune

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if len(stack) > 0 && brackets[char] == stack[len(stack)-1] {
			//删除尾部一个元素
			stack = stack[:len(stack)-1]
		} else {
			return false
		}

	}
	return len(stack) == 0
}
