package leetcode

func truncateSentence(s string, k int) string {
	res := ""
	c := 0
	for i := range s {
		if s[i] == ' ' {
			c++
		}
		if c == k {
			break
		}
		res = res + string(s[i])
	}
	return res
}
