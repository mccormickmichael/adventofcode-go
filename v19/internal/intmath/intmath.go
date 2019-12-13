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

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return Gcd(b % a, a)
}

func Lcm(a, b int) int {
	return  a * b / Gcd(a, b)
}