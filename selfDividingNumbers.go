package leetcode

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if isSelfDiv(i) {
			res = (append(res, i))
		}
	}
	return res
}

func isSelfDiv(n int) bool {
	res := true
	s := strconv.Itoa(n)
	for i := range s {
		num, _ := strconv.Atoi(string(s[i]))
		if num == 0 {
			res = false
			break
		} else if n%num != 0 {
			res = false
			break
		}

	}
	return res
}
