package test

func EqualIntSlice(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, v := range lhs {
		if v != rhs[i] {
			return false
		}
	}
	return true
}
