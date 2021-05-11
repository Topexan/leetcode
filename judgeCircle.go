package leetcode

func judgeCircle(moves string) bool {
	xCoor := 0
	yCoor := 0
	res := false
	for _, i := range moves {
		switch i {
		case 'L':
			xCoor--
		case 'R':
			xCoor++
		case 'U':
			yCoor++
		case 'D':
			yCoor--
		}
	}
	if xCoor == 0 && yCoor == 0 {
		res = true
	}
	return res
}
