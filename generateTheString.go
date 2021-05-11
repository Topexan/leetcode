package leetcode

func generateTheString(n int) string {
	res := ""
	if n%2 == 0 {
		if (n/2)%2 == 0 {
			for i := 0; i < n; i++ {
				if i < (n/2)+1 {
					res = res + "x"
				} else {
					res = res + "y"
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if i < n/2 {
					res = res + "x"
				} else {
					res = res + "y"
				}
			}
		}
	} else {
		res = generateTheString(n-1) + "z"
	}
	return res
}
