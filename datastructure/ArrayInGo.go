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

// ValueType 值类型
func ValueType() {
	var arrayA [5]int = [5]int{1, 2, 3, 4, 5}
	arrayB := arrayA
	fmt.Println("这里我们定义了两个数组: arrayA 和 arrayB，其中arrayA是有初始值的，而arrayB := arrayA")
	fmt.Printf("arrayA: %v, arrayB: %v\n", arrayA, arrayB)
	fmt.Println("这里我们将arrayB[0]中的值从1改为11，这个时候我们看一下两个数组中的值")
	fmt.Printf("arrayA: %v, arrayB: %v\n", arrayA, arrayB)
	fmt.Println("我们可以看出 arrayB[0]的值发生了改变，而arrayA[0]中的值还是保持不变")
	fmt.Println("这就是我们说的值传递，即：赋值和传参会复制整个数组，因此改变副本的值，不会改变本身的值")
	fmt.Println("和数组相似的数据类型切片就是引用传递")
	sliceA := arrayA[1:4]
	fmt.Printf("我们这里定义了一个arrayA：%v的切片sliceA:%v, 接下来我们将sliceA[0]中的值改为5\n", arrayA, sliceA)
	sliceA[0] = 5
	fmt.Printf("我们分别查看一下arrayA: %v, sliceA: %v.\n", arrayA, sliceA)
	fmt.Println("我们可以看到值改变了sliceA中的值对应的arrayA的值也发生了改变。这就是引用传递，即：改变副本的值，会改变本身的值")
}

// MultidimensionalArray 多维数组
func MultidimensionalArray() {
	var marray = [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	fmt.Println(marray)
	fmt.Printf("多维数组的取值；如果我们想拿到数组中第二个数组的第2个值，那么我们可以这么写marry[1][1]: %v\n", marray[1][1])
}
