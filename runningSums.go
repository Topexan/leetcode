package Leetcode

func runningSum(nums []int) []int {
	var res []int
	c := 0
	for i := range nums {
		c = c + nums[i]
		res = append(res, c)
	}
	return res
}
