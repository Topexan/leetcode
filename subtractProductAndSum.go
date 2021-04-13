package Leetcode

func subtractProductAndSum(n int) int {
	a := 1
	b := 0
	for n > 0 {
		a = a * (n % 10)
		b = b + (n % 10)
		n = n / 10
	}
	return a - b
}
