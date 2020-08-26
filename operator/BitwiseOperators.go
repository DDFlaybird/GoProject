package operator
/*
	& 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 两位都是1才为1
	| 按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或。 两位有一个为1就是1
	^ 按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或。 两位不同才为1
	<< 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
	>> 	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。
*/

import "fmt"

// BitwiseOperatorTestCase 展示位运算符
func BitwiseOperatorTestCase() {
	var a = 5
	var b = 2
	var c = 1
	fmt.Println("***************BitwiseOperatorTestCase*****************")
	fmt.Printf("定义变量a:%d    二进制值为：%b\n", a, a)
	fmt.Printf("定义变量b:%d    二进制值为：%b\n", b, b)
	fmt.Printf("定义变量c:%d    二进制值为：%b\n", c, c)
	fmt.Println("a&b=", a&b)
	fmt.Println("a|b=", a|b)
	fmt.Println("a^b=", a^b)
	fmt.Println("a>>2=", c>>2)
	fmt.Println("a<<2=", c<<2)
	fmt.Println("***************BitwiseOperatorTestCase*****************")
}
