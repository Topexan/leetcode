package Leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var res []bool
	var num int
	max := 0
	for i := range candies {
		num = candies[i] + extraCandies
		for j := range candies {
			if max < candies[j] {
				max = candies[j]
			}
		}
		if max <= num {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
		max = 0
	}
	return res
}
