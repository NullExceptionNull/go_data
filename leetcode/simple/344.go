package simple

func reverseString(s []byte) []byte {
	var right = len(s) - 1

	for left := 0; left < right; left++ {
		c := s[left]
		s[left] = s[right]
		s[right] = c
		right--
	}
	return s
}
