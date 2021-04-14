package Leetcode

func maxDepth(s string) int {
	res := 0
	num := 0
	for i := range s {
		if s[i] == '(' {
			num++
			if num > res {
				res = num
			}
		} else if s[i] == ')' {
			num--
		}
	}
	return res
}
