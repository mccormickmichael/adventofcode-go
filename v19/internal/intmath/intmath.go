package intmath

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Cmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}


