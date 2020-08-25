/*
str主要用来解决类型不同类型转换的库。
同类型转换如(int8 -> int16)可以直接使用int16(int8);但是涉及到不同类型的转换如int8 -> string的时候需要用strconv这个库。
tips:
fmt这个库也可以实现类型转换但是只是将各种类型转换成string。但是不能将string类型转换成别的格式。
*/
package typecoversion

import (
	"fmt"
	"strconv"
)

func StrCovert(inputInt string) int64 {
	returnInt, _ := strconv.ParseInt(inputInt, 10, 64)
	return returnInt
}

// fmt type conversion example
// 使用Sprintf来将各个类型的格式转换成string格式
func FmtIntConvertExample(inputValue int) string {
	return fmt.Sprintf("%d", inputValue)
}

func FmtByteConvertExample(inputValue byte) string {
	return fmt.Sprintf("%c", inputValue)
}

func FmtBoolConvertExample(inputValue bool) string {
	return fmt.Sprintf("%t", inputValue)
}
