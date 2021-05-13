package leetcode

func finalPrices(prices []int) []int {
	var c int
	for i := 0; i < len(prices)-1; i++ {
		c = 0
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				c = prices[j]
				break
			}
		}
		prices[i] -= c
	}
	return prices
}
