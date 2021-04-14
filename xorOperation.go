package Leetcode

func xorOperation(n int, start int) int {
	res := start
	if n == 1 {
		return start
	} else {
		for i := 0; i < n-1; i++ {
			start = start + 2
			res = res ^ start

		}
	}
	return res
}
