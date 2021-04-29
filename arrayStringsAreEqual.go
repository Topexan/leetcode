package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	s2 := ""
	var res bool
	for i := range word1 {
		s1 = s1 + word1[i]
	}
	for j := range word2 {
		s2 = s2 + word2[j]
	}
	if s1 == s2 {
		res = true
	} else {
		res = false
	}
	return res
}
