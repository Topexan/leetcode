package leetcode

import "strings"

func reverseWords(s string) string {
	res := ""
	arr := strings.Split(s, " ")
	for i := range arr {
		for j := len(arr[i]) - 1; j >= 0; j-- {
			res += string(arr[i][j])
		}
		if i != len(arr)-1 {
			res += " "
		}
	}
	return res
}
