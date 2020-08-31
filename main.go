package main

import (
	"GoProject/datastructure"
	"fmt"
)

func main() {
	fmt.Println("***************typecoversion*****************")
	// ac := 3
	// bc := 10
	// a := "100"
	// b := 10000000
	// // 类型转换
	// fmt.Println(typecoversions.StrCovert(a))
	// fmt.Println(typecoversions.FmtIntConvertExample(b))
	// fmt.Println(bc % ac)
	fmt.Println("***************typecoversion*****************")

	// 操作符
	// operator.BitwiseOperatorTestCase()

	// 流程控制
	fmt.Println("***************FLowControl*****************")
	// flowcontrol.AnthorSimpleIfExample(22)
	// flowcontrol.ElifSimpleExample(52)
	// flowcontrol.ForSimpleExample()
	// flowcontrol.EventNumberExample(22)
	// flowcontrol.SumExample(12)
	// flowcontrol.TimesRangeExample(100, 9)
	// flowcontrol.FactorialExample(5)
	// flowcontrol.RectangleExample(10, 20)
	// flowcontrol.PrintTriangleExample(20, 10)
	// flowcontrol.MultiplicationTableExample(9)
	// flowcontrol.SwitchExample("E")
	// flowcontrol.SwitchCaseExampleB()
	// flowcontrol.SwitchCaseExampleC(60)
	// flowcontrol.BreakInFor()
	// flowcontrol.BreakWithLable()
	// flowcontrol.ContinueWithLable()
	// flowcontrol.GotoWithLable()
	fmt.Println("***************FLowControl*****************")
	fmt.Println()
	fmt.Println("***************DataStructure*****************")
	// datastructure.RangeWithArray()
	// datastructure.ForWithArray()
	var arraySample = [...]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	var findsumandindex = []int{1, 2, 3, 4, 5}
	datastructure.ArraySumAndAverage(arraySample)
	datastructure.FindSumAndIndex(findsumandindex, 8)
	datastructure.MaxAndMinIntInArray(findsumandindex)
	fmt.Println("***************DataStructure*****************")
}
