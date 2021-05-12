package leetcode

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	var c int
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				c = nums[j]
				nums[j] = nums[i]
				nums[i] = c
			}
		}
	}
	return nums
}
