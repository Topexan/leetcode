func smallerNumbersThanCurrent(nums []int) []int {
    var res []int
    l := len(nums)
    s := 0
    for i := range nums {
        for j := 0; j < l; j++ {
            if nums[j] < nums[i] {
                s++
            }
        }
        res = append(res, s)
        s = 0
    }
    return res
}