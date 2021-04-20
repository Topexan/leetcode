package Leetcode

func minOperations(boxes string) []int {
	res := make([]int, len(boxes))

	for i := range boxes {
		for j := range boxes {
			if boxes[j] == '1' && j != i {
				res[i] += abs(i - j)
			}
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
