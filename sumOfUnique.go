package leetcode

func sumOfUnique(nums []int) int {
	res := 0
	for i := range nums {
		c := 0
		for j := range nums {
			if nums[i] == nums[j] {
				c++
			}
		}
		if c == 1 {
			res += nums[i]
		}
	}
	return res
}
