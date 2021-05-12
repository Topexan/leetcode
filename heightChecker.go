package leetcode

func heightChecker(heights []int) int {
	res := 0
	var c int
	var h2 []int
	for i := range heights {
		h2 = append(h2, heights[i])
	}

	for i := 0; i < len(heights)-1; i++ {
		for j := i; j < len(heights); j++ {
			if heights[i] > heights[j] {
				c = heights[j]
				heights[j] = heights[i]
				heights[i] = c
			}
		}
	}
	for i := range h2 {
		if h2[i] != heights[i] {
			res++
		}
	}
	return res
}
