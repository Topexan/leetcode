func isPalindrome(x int) bool {
    s1 := strconv.Itoa(x)
    s2 := ""
    for i := len(s1) - 1; i >= 0; i-- {
        s2 = s2 + string(s1[i])
    }
    if (x < 0) {
        return false
    }
    if (s1 != s2) {
        return false
    }
    return true
}