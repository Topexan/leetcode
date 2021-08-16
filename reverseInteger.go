func reverse(x int) int {
    res := 0
    checkMinus := false
    if (x > 2147483648) || (x < -2147483648) {
        return 0
    }
    if x < 0 {
        x = x * -1
        checkMinus = true
    }
    for x > 0 {
        res = (res * 10) + (x % 10)
        x /=10
    }
    if (checkMinus) {
        res = res * -1
    }
    if (res > 2147483648) || (res < -2147483648) {
        return 0
    }
    return res
}