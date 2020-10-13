package datastructure

import "fmt"

// 自定义类型
type MyInt int64

// 给自定义类型绑定方法
func (m MyInt) ShowInfo() {
	fmt.Println(m)
}

// 自定义函数 function_in_go中第55，56行有实际应用
type myFn func(int, int) int

// 类型别名
type myFloat = float32

type MyStruct struct {
	name string
	age  int8
}

func CustomType() {
	var a MyInt = 20
	fmt.Printf("%v, %T\n", a, a) // 20, datastructure.myInt
	// 类型别名和自定义类型的区别
	// 主要差别在类别上面一个类型。
	// 类型别名本质上还是float32或者你指定的类型
	var b myFloat = 3.14
	fmt.Printf("%v, %T\n", b, b) // 3.14, float32
}

func StructInit() {
	// 1 其中一种实例化方式
	var TestStruct MyStruct
	TestStruct.name = "string"
	TestStruct.age = 22
	fmt.Println(TestStruct)
	// 2 第二种 使用new方法实例化
	var test2 = new(MyStruct)
	test2.name = "jason"
	test2.age = 23
	// 使用指针更改结构体类型
	(*test2).name = "jack"
	fmt.Printf("值: %#v, 类型：%T\n", test2, test2)
	// 3 第三种直接使用&符号来实例化 推荐
	test3 := &MyStruct{
		name: "luna",
		age:  22,
	}
	fmt.Printf("值: %#v, 类型：%T\n", test3, test3)

	// 4 另外一种使用&来实例化结构体的方法
	var test4 = &MyStruct{}
	test4.name = "mafei"
	test4.age = 25
	fmt.Printf("值: %#v, 类型：%T\n", test4, test4)
	// 5 键值对实例化 推荐
	test5 := MyStruct{
		name: "jackson",
		age:  22,
	}
	fmt.Printf("值: %#v, 类型：%T\n", test5, test5)
	var test6 = &MyStruct{
		name: "huck",
		age:  22,
	}
	fmt.Printf("值: %#v, 类型：%T\n", test6, test6)

	test7 := MyStruct{
		name: "bart",
	}
	fmt.Printf("值: %#v, 类型：%T\n", test7, test7)

	test8 := MyStruct{
		"lisa",
		22,
	}
	// 列表传值，不使用键值对
	fmt.Printf("值: %#v, 类型：%T\n", test8, test8)
	// 结构体是值类型，
	test9 := test8
	test9.name = "bart"
	fmt.Printf("值: %#v, 类型：%T\n", test8, test8) // 值: datastructure.MyStruct{name:"lisa", age:22}, 类型：datastructure.MyStruct
	// 改变9的值并不会对8的值产生什么影响
}

// 通过结构体来实现oop中的类
// 结构体方法
/*
格式：
	func(接受者变量 接收者类型)方法名(参数列表)(返回值列表){
		函数体
	}
*/
// 接收者函数 接收者类似python中的self和js中的this
type Task struct {
	TotalTask     int
	TotalWorkTime float32
}

func (t Task) ShowTotalTask() {
	// t 就是接收者
	fmt.Println(t.TotalTask)
}

// 这个函数并不会修改实例中阿totaltask和totalworktime，
// 函数体内之执行的函数只是修改t这个Task实列的的，并不能将传入进来的实例的属性修改
// 如果你想哟修改那么就要传入一个执政类型
func (t Task) Edit(totalTask int, totalWorkTime float32) {
	t.TotalTask = totalTask
	t.TotalWorkTime = totalWorkTime
}

// 这个就会修改传进来struct实列的属性
func (t Task) Edit2(totalTask int, totalWorkTime float32) {
	t.TotalTask = totalTask
	t.TotalWorkTime = totalWorkTime
}
