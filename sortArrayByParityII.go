package leetcode

func sortArrayByParityII(nums []int) []int {
	evenI := 0
	oddI := 1
	res := make([]int, len(nums))
	for i := range nums {
		if nums[i]%2 == 0 {
			res[evenI] = nums[i]
			evenI += 2
		} else {
			res[oddI] = nums[i]
			oddI += 2
		}
	}
	return res
}
