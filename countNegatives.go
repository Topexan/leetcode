package leetcode

func countNegatives(grid [][]int) int {
    res := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] < 0 {
                res++
            }
        }
    }
    return res
}