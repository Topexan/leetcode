package leetcode

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	var arr []int
	for i := range points {
		arr = append(arr, points[i][0])
	}
	sort.Ints(arr)
	res := 0
	for i := len(arr) - 1; i > 0; i-- {
		if (arr[i] - arr[i-1]) > res {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}
