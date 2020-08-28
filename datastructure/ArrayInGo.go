package datastructure

import "fmt"

/*
	四种初始化数组的方式
	一、
		var arraysample [3]int
		arraysample[0] = 1
		arraysample[1] = 2
		arraysample[2] = 3
	二、
		arraysample := [3]int{1, 2, 3}
		var arraysample = [3]int{1, 2, 3}
	三、
		[...]自动推导的意思
		arraysample := [...]int{1, 2, 3, 4, 5}
	四、
		arraysample := [...]int{1:2, 0:1, 2:3}
		指定值在数组中的位置
*/

// RangeWithArray 数组遍历
func RangeWithArray() {
	arraysample := [5]int{1, 2, 3, 4, 5}
	for i, v := range arraysample {
		fmt.Printf("index:%d, value:%d\n", i, v)
	}

}

// ForWithArray for循环遍历数组
func ForWithArray() {
	var arraysample [10]int
	for i := 0; i < len(arraysample); i++ {
		fmt.Println(arraysample[i])
	}
}

// ArraySumAndAverage 返回一个平均值和总和
func ArraySumAndAverage(array [12]int) (int, float64) {
	var sum int
	var num int = len(array)
	for _, v := range array {
		sum += v
	}
	average := float64(sum) / float64(num)
	fmt.Printf("sum: %d, average: %.4f\n", sum, average)
	return num, average
}


// 