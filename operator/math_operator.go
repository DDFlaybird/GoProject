package operator

// Add 相加
func Add(x, y int) int {
	return x + y
}

// Minus 相减
func Minus(x, y int) int {
	return x - y
}

// Divided 除以
func Divided(x, y int) int {
	return x / y
}

// Mul 乘法
func Mul(x, y int) int {
	return x * y
}

// Remainder 求余数
/*
eg:
	x = 3
	y = 10
余数：
	10 - (10 / 3 后向下取整) * 3
	10 - 3*3 = 1
向下取整：
	3.14 向下取整 3
	6.7546 向下取整 6
	-3.14 向下取整 -4
*/
func Remainder(x, y int) int {
	return x % y
}

// AutoIncrement 自增
// 自增是单独语句并不是运算符
func AutoIncrement(x int) int {
	x++
	return x
}

// Decrementing 自减
// 自减是单独语句并不是运算符
func Decrementing(x int) int {
	x--
	return x
}
