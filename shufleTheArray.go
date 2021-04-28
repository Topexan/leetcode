package leetcode

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	c := 0
	for i := 0; i < len(nums)/2; i++ {
		res[c] = nums[i]
		c += 2
	}
	c = 1
	for j := (len(nums) / 2); j < len(nums); j++ {
		res[c] = nums[j]
		c += 2
	}
	return res
}
