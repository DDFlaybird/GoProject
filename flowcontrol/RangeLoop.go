package flowcontrol

import "fmt"

// RangeString 用for-range遍历字符串
// k是索引， v是值
// 一个汉字字符为3个字节所以k可能不连贯
func RangeString(str string) {
	for k, v := range str {
		fmt.Printf("key=%v, value=%c\n", k, v)
	}
}

// RangeSlice 用for-range遍历切片
func RangeSlice(slices []int){
	for k, v := range slices{
		fmt.Printf("key=%v, value=%c\n", k, v)
	}
}