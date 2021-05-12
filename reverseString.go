package leetcode

func reverseString(s []byte) {
	l := len(s) - 1
	var c byte
	for i := 0; i < len(s)/2; i++ {
		c = s[l]
		s[l] = s[i]
		s[i] = c
		l--
	}
}
