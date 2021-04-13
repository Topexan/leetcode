package Leetcode

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))
	for i := range indices {
		b[indices[i]] = s[i]
	}
	return string(b)
}
