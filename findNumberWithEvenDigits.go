package leetcode

func findNumbers(nums []int) int {
	res := 0
	for i := range nums {
		if digits(nums[i])%2 == 0 {
			res++
		}
	}
	return res
}

func digits(n int) int {
	res := 0
	for i := n; i > 0; i /= 10 {
		res++
	}
	return res
}
