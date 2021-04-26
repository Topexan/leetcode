package leetcode

import "sort"

func maxIceCream(costs []int, coins int) int {
	res := 0
	sort.Ints(costs)

	for i := range costs {
		if costs[i] <= coins {
			coins = coins - costs[i]
			res++
		}
	}
	return res
}
