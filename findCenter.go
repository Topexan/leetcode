package leetcode

func findCenter(edges [][]int) int {
	num1 := edges[0][0]
	num2 := edges[0][1]
	c1 := 0
	c2 := 0
	for i := range edges {
		for j := range edges[i] {
			if edges[i][j] == num1 {
				c1++
			} else if edges[i][j] == num2 {
				c2++
			}
		}
	}
	if c1 > c2 {
		return num1
	}
	return num2
}
