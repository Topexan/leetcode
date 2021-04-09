package Leetcode

func defangIPaddr(address string) string {
	res := ""
	for i := range address {
		if address[i] == '.' {
			res = res + "[.]"
		} else {
			res = res + string(address[i])
		}
	}
	return res
}
