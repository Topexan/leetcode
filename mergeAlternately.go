package leetcode

func mergeAlternately(word1 string, word2 string) string {
	var l int
	res := ""
	if len(word1) > len(word2) {
		l = len(word1)
	} else {
		l = len(word2)
	}
	for i := 0; i < l; i++ {
		if i < len(word1) {
			res = res + string(word1[i])
		}
		if i < len(word2) {
			res = res + string(word2[i])
		}
	}
	return res
}
