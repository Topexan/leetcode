package Leetcode

func balancedStringSplit(s string) int {
	res := 0
	r := 0
	l := 0
	for i := range s {
		if s[i] == 'R' {
			r++
		}
		if s[i] == 'L' {
			l++
		}
		if (r == l) && r != 0 {
			res++
			r = 0
			l = 0
		}
	}
	return res
}
