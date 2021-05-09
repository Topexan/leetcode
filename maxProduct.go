func maxProduct(nums []int) int {
    max := 0
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            n := (nums[i] - 1) * (nums[j] - 1)
            if max < n {
                max = n
            }
        }
    }
    return max
}