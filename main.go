package main

import (
	"GoProject/datastructure"
	"GoProject/functions"
	"fmt"
)

func test(strings string) string {
	return strings + "test"
}

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
	mapExample := map[int]int{45: 23, 6: 2, 7: 3, 4: 5, 222: 5}
	datastructure.ArraySumAndAverage(arraySample)
	datastructure.FindSumAndIndex(findsumandindex, 8)
	datastructure.MaxAndMinIntInArray(findsumandindex)
	datastructure.ValueType()
	datastructure.MultidimensionalArray()
	datastructure.CapOfSlice()
	datastructure.AppendInSlice()
	datastructure.CopyInSlice()
	datastructure.DeleteInSlice()
	datastructure.ForRangeInMap()
	datastructure.CURDInMap()
	datastructure.IntoMap()
	datastructure.SortMap(mapExample)
	fmt.Println("***************DataStructure*****************")
	fmt.Println()
	fmt.Println("***************Functions*****************")
	functions.FlexableArgs(1, 2, 3, 4, 5)
	functions.NameReturnValue(1, 2)
	functions.FuncType()
	functions.FuncArgsTwo(1, 1, func(x, y int) int { return x - y })
	var returnfunc = functions.FuncReturn("+")
	fmt.Println(returnfunc(1, 2))
	// 匿名函数可在函数内部调用
	func() {
		fmt.Println("匿名函数的调用")
	}()
	fmt.Println(functions.RecursionFunc(4))
	// 闭包
	functions.FuncCloser()
	functions.DeferExample()
	functions.LambdaReturn()
	c := functions.C()
	d := functions.D()
	fmt.Println("A return:", functions.A()) // 打印结果为 A return: 0
	fmt.Println("B return:", functions.B()) // 打印结果为 A return: 0
	fmt.Println("c return:", c, c)          // 打印结果为 c return: 2 0xc082008340
	fmt.Println("d return:", d, &d)         // 打印结果为 c return: 2 0xc082008340
	fmt.Println(functions.DeferFunc(1))
	fmt.Println("***************Functions*****************")
}
