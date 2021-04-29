package leetcode

func numberOfMatches(n int) int {
    res := 0
    for n > 1 {
        if n % 2 == 0 {
            n /= 2
            res += n
        } else {
            n = (n - 1) / 2
            res += n
            n++
        }
    }
    return res
}