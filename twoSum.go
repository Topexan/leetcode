package Leetcode

func twoSum(nums []int, target int) []int {
    var res []int
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                res = append(res, i)
                res = append(res, j)
                break
            }
        }
    }
    return res
}