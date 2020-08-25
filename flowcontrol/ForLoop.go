package flowcontrol

import "fmt"

/*
	for 初始语句；条件表达式； 结束语句{
			循环体语句
	}
	初始语句和结束语句都可以省略。
*/

// ForSimpleExample 简单for循环解释
func ForSimpleExample() {
	fmt.Println("For")
	for x := 2; x < 10; x++ {
		fmt.Println(x)
	}
}

// WhileSimpleExample GO中的while
// 单独一个for关键字就是无限循环， 使用while内部一定有跳出语句
func WhileSimpleExample() {
	fmt.Println("While")
	var i int8 = 1
	for {
		if i < 100 {
			fmt.Println(i)
			i++
		} else {
			break
		}
	}
}

// EventNumberExample while语句的联系
func EventNumberExample(rangeNumber int) []int {
	fmt.Printf("输出在%d范围内的偶数\n", rangeNumber)
	var eventNumber []int
	for i := 0; i < rangeNumber; i++ {
		if i%2 == 0 {
			eventNumber = append(eventNumber, i)
		}
	}
	fmt.Println(eventNumber)
	return eventNumber
}

// SumExample 递增求和
func SumExample(rangeNUmber int) int {
	var sum int = 0
	for i := 0; i <= rangeNUmber; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum
}

// TimesRangeExample 打印1~100之间所有是9的倍数的个数及总和
func TimesRangeExample(rangeNUmber, time int) (int, int) {
	var num int = 0
	var sum int = 0
	for i := 0; i <= rangeNUmber; i++ {
		if i%time == 0 {
			fmt.Println(i)
			num += 1
			sum += i
		}
	}
	fmt.Printf("在%d范围之内有%d个%d的倍数\n", rangeNUmber, num, time)
	fmt.Printf("总和为:%d\n", sum)
	return sum, num
}
