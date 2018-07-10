//返回这个数值应该在的位置
func binarySearch(base []int, x int) int {
	down, up := 0, len(base)-1
	if x <= base[0] {
		return 0
	} else if x >= base[up] {
		return up + 1
	}

	for {
		if down+1 == up {
			return up
		}
		bidx := (down + up) / 2
		switch {
		case x < base[bidx]:
			up = bidx
		case x > base[bidx]:
			down = bidx
		default:
			return bidx
		}
	}
}
