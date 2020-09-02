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

/*切片*/

/*
	初始化切片的两种方式：
		1、var a []int
		2、a := make([]int, 2, 5) ; make函数的详解 make([]type, len, cap) []type 切片中元素内容，len长度，cap容量
	切片的初始化：
		a = []int{1, 2, 3, 4}
		a = append(a, 1); append2倍扩容
	未初始化的切片是nil 和数组不同；数组中如[2]string 如果没有初始化那么数组中是""，其他类型的数组就是对应类型的初始值
	切片的长度和容量：
	长度：切片的长度就是他包含的个数
	容量：切片的容量是从他的第一个元素开始数，到其底层数组元素末尾的个数
	容量例子：
		a := []int{1, 2, 3, 4, 5, 6, 6}
		a[1:] // [2， 3， 4， 5， 6， 6] 这个长度就是6, 容量也是6
		a[1:4] // [2, 3, 4] 这个长度是3，容量是6 为啥呢？ 重点就是这句话 ”到其底层数组元素的个数“ 也就是说他的容量就是2 - 6 之间的容量
*/

// CapOfSlice 解释容量
func CapOfSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("切片：%v, 长度：%d, 容量：%d\n", a, len(a), cap(a))
	sliceA := a[1:4]
	fmt.Printf("定义一个数组的切片：%v, 长度：%d, 容量：%d\n", sliceA, len(sliceA), cap(sliceA))
	fmt.Println("这个时候cap和长度并不相等。那么为啥呢？记住cap的定义：“切片的容量是从他的第一个元素开始数，到其底层数组元素末尾的个数”，所以容量是从2开始数到元素组的末尾元素也就是6结束")
}

// AppendInSlice append函数在切片中的使用
// append 可以给切片扩容
// append 可以给切片添加多个值
// append 可以合并切片
/*
切片的扩容策略：
			当需要的容量超过原切片容量的两倍时，会使用需要的容量作为新容量；
			当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量。
*/
func AppendInSlice() {
	var language []string
	language = append(language, "Go")
	// append 可以给切片进行扩容
	fmt.Printf("切片：%v, 长度：%d, 容量：%d\n", language, len(language), cap(language))
	language = append(language, "Python")
	fmt.Printf("切片：%v, 长度：%d, 容量：%d\n", language, len(language), cap(language))
	var languageA []string
	// 使用append合并切片
	languageA = append(languageA, "Perl", "JavaScript")
	language = append(language, languageA...)
	fmt.Println(language)
}

// CopyInSlice copy
func CopyInSlice() {
	// copy的时候要使用make并且指明容量和长度
	sliceB := make([]int, 2, 16)
	sliceA := []int{1, 2, 3, 4}
	copy(sliceB, sliceA)
	fmt.Println(sliceB)
}

// DeleteInSlice 删除
// slice中没有删除的接口所以只能通过append合并数组的方式删除切片中的元素
func DeleteInSlice() {
	slices := []int{1, 2, 3, 4, 5, 6}
	slices = append(slices[0:1], slices[2:]...)
	fmt.Println(slices)
	// append 只能应用在slice中
	// 所以我们通过引用传递的机制 对数组切片然后再对切片appen删除需要 删除的数据，但是由于数组的底层原理，我们删除元素后数组的长度不会改变，所以他会给我们补位一个元素。
	arrays := [6]int{1, 2, 3, 4, 5, 6}
	sliceB := arrays[0:]
	sliceB = append(sliceB[0:1], sliceB[2:]...)
	fmt.Println(arrays)
}