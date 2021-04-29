package leetcode

func toLowerCase(str string) string {
    res := ""
    for i := range str {
        if str[i] > 64 && str[i] < 91 {
            res = res + string(str[i] + 32)
        } else {
            res = res + string(str[i])
        }
    }
    return res
}