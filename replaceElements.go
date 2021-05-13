package leetcode

func replaceElements(arr []int) []int {
	var res []int
	if len(arr) == 1 {
		return []int{-1}
	}
	var max int
	for i := 0; i < len(arr)-1; i++ {
		max = 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		res = append(res, max)
	}
	res = append(res, -1)
	return res
}
