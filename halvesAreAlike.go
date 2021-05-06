package leetcode

func halvesAreAlike(s string) bool {
	res := false
	count1 := 0
	count2 := 0
	for i := range s {
		if i < len(s)/2 {
			switch s[i] {
			case 'a':
				count1++
			case 'e':
				count1++
			case 'i':
				count1++
			case 'o':
				count1++
			case 'u':
				count1++
			case 'A':
				count1++
			case 'E':
				count1++
			case 'I':
				count1++
			case 'O':
				count1++
			case 'U':
				count1++
			}
		} else {
			switch s[i] {
			case 'a':
				count2++
			case 'e':
				count2++
			case 'i':
				count2++
			case 'o':
				count2++
			case 'u':
				count2++
			case 'A':
				count2++
			case 'E':
				count2++
			case 'I':
				count2++
			case 'O':
				count2++
			case 'U':
				count2++
			}

		}
	}
	if count1 == count2 {
		res = true
	}
	return res
}
