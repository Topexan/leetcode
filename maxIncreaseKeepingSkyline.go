package leetcode

func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	var tb, lr []int
	var minTB int
	var minLR int
	for i := 0; i < len(grid); i++ {
		minTB = 0
		minLR = 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > minLR {
				minLR = grid[i][j]
			}
			if grid[j][i] > minTB {
				minTB = grid[j][i]
			}
		}
		lr = append(lr, minLR)
		tb = append(tb, minTB)
	}
	h := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if tb[i] < lr[j] {
				h = tb[i]
			} else {
				h = lr[j]
			}
			res = res + (h - grid[i][j])
		}
	}
	return res
}
