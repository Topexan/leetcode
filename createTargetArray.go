package Leetcode

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i, pos := range index {
		copy(res[pos+1:i+1], res[pos:i])
		res[pos] = nums[i]
	}
	return res
}
