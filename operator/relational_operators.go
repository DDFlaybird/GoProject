package operator

// Equal 等于
func Equal(x, y int) bool {
	return x == y
}

// Bigger 大于
func Bigger(x, y int) bool {
	return x > y
}

// Ge 大于等于
func Ge(x, y int) bool {
	return x >= y
}

// LessThan 小于
func LessThan(x, y int) bool {
	return x < y
}

// Le 小于等于
func Le(x, y int) bool {
	return x <= y
}

// Neq 不等于
func Neq(x, y int) bool {
	return x != y
}
