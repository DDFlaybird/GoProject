package main

import (
	"GoProject/FlowControl"
	"GoProject/Operator"
	"GoProject/TypeCoversion"
	"fmt"
)

func main() {
	fmt.Println("***************TypeCoversion*****************")
	ac := 3
	bc := 10
	a := "100"
	b := 10000000
	// 类型转换
	fmt.Println(TypeCoversion.StrCovert(a))
	fmt.Println(TypeCoversion.FmtIntConvertExample(b))
	fmt.Println(bc % ac)
	fmt.Println("***************TypeCoversion*****************")
	// 操作符
	Operator.BitwiseOperatorTestCase()
	// 流程控制
	fmt.Println("***************FLowControl*****************")
	FlowControl.AnthorSimpleIfExample(22)
	FlowControl.ElifSimpleExample(52)
	//FlowControl.ForSimpleExample()
	FlowControl.EventNumberExample(22)
	FlowControl.SumExample(12)
	FlowControl.TimesRangeExample(100, 9)
	fmt.Println("***************FLowControl*****************")
}
