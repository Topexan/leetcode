package leetcode

func squareIsWhite(coordinates string) bool {
    var res bool
    if (coordinates[0] % 2 == coordinates[1] % 2) {
        res = false
    } else {
        res = true
    }
    return res
}