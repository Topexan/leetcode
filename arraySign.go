package leetcode

func arraySign(nums []int) int {
    minus := 0
    res := 1
    for _, i := range nums {
        if i == 0 {
            res = 0
            break
        }
        if i < 0 {
            minus++
        }
    }
    if minus % 2 != 0 {
        res = -1
    }
    return res
}