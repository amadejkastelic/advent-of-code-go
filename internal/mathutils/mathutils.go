package mathutils

func Mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
