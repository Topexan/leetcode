package leetcode

func repeatedNTimes(A []int) int {
	var res int
	n := len(A) / 2
	for i := range A {
		r := 0
		for j := range A {
			if A[i] == A[j] {
				r++
			}
		}
		if r == n {
			res = A[i]
			break
		}
	}
	return res
}
