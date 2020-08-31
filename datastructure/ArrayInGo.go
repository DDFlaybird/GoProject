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

// FindSumAndIndex 从数组中找到和为指定数的两个元素的值以及他们在数组中的位置。
// TODO 算法待优化
func FindSumAndIndex(array []int, result int) []map[string][]int {
	var indexAndresult []map[string][]int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if (array[i] + array[j]) == result {
				indexs := map[string][]int{"element": []int{array[i], array[j]}, "index": []int{i, j}}
				indexAndresult = append(indexAndresult, indexs)
			}
		}
	}
	fmt.Println(indexAndresult)
	return indexAndresult
}

// MaxAndMinIntInArray 求数组中最大值
func MaxAndMinIntInArray(arrays []int) (int, int) {
	var min = arrays[0]
	var max = arrays[0]
	for _, v := range arrays {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Printf("最大值：%d, 最小值：%d\n", max, min)
	return max, min
}
