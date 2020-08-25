package main

import (
	"GoProject/flowcontrol"
	"GoProject/operator"
	"GoProject/typecoversion"
	"fmt"
)

func main() {
	fmt.Println("***************TypeCoversion*****************")
	ac := 3
	bc := 10
	a := "100"
	b := 10000000
	// 类型转换
	fmt.Println(typecoversion.StrCovert(a))
	fmt.Println(typecoversion.FmtIntConvertExample(b))
	fmt.Println(bc % ac)
	fmt.Println("***************TypeCoversion*****************")
	// 操作符
	operator.BitwiseOperatorTestCase()
	// 流程控制
	fmt.Println("***************FLowControl*****************")
	flowcontrol.AnthorSimpleIfExample(22)
	flowcontrol.ElifSimpleExample(52)
	//FlowControl.ForSimpleExample()
	flowcontrol.EventNumberExample(22)
	flowcontrol.SumExample(12)
	flowcontrol.TimesRangeExample(100, 9)
	fmt.Println("***************FLowControl*****************")
}
