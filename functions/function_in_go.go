package functions

import (
	"fmt"
	"reflect"
)

/*
全局变量和局部变量
和python中的定义一样
全局变量：定义在函数外部，整个程序运行周期之间都有效
局部变量：定义在函数内部，只能函数内部可以调用，函数外不能调用(闭包)
*/

// FlexableArg 可变参数
// 可变参数为一个列表
func FlexableArg(x ...int) {
	fmt.Printf("参数：%v, 类型是：%v\n", x, reflect.TypeOf(x))
}

// FlexableArgs 函数中可变参数和一般参数
func FlexableArgs(x int, y ...int) {
	fmt.Printf("参数：%v, 类型是：%v\n", x, reflect.TypeOf(x))
	fmt.Printf("参数：%v, 类型是：%v\n", y, reflect.TypeOf(y))
}

// NameReturnValue 返回值命名
// 返回值命名就是提前声明返回值变量，
// 这样再函数中就不需要特意写出来。一种便捷的方式。
// 效果等价与再函数内 var sum int var sub int
// 或者 sum := x + y; sub := x - y 列表推导
func NameReturnValue(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x - y
	sub = x + y
	return sum, sub
}

// 函数类型和名称

// Stringfunc 任意函数如果满足函数体定义的格式都可以将这个函数赋值给声明的类型
// 自定义类型
type Stringfunc func(string) string

func hello(name string) string {
	return name + "，hello！"
}

func goodbye(name string) string {
	return "goodbye, " + name + "."
}

// FuncType 函数类型
func FuncType() {
	var types Stringfunc
	types = hello
	fmt.Println(types("leo"))
}

// myfunctype 自定义类型
type myfunctype func(int, int) int

// FuncArgs 函数作为参数
func FuncArgs(x, y int, method func(int, int) int) {
	fmt.Println("函数作为参数的一种实现方式")
}

// FuncArgsTwo 函数作为参数的另外一种实现方式
func FuncArgsTwo(x, y int, method myfunctype) {
	method(x, y)
	fmt.Println("另外一种函数作为参数的实现方式")
}

// FuncReturn 函数作为返回值
func FuncReturn(mark string) myfunctype {
	switch mark {
	case "+":
		return func(x, y int) int { return x + y }
	case "-":
		return func(x, y int) int { return x - y }

	case "*":
		return func(x, y int) int { return x * y }

	case "/":
		return func(x, y int) int { return x / y }
	default:
		return nil
	}
}

/*
匿名函数：
	匿名韩式的格式 func (参数) 返回格式 {}
	普通方法不可以放到另一个方法内部调用，但是匿名方法可以
*/

/*
函数的递归调用
递归就是函数自己调用自己
*/
// RecursionFunc 递归求得1-100的和
func RecursionFunc(x int) int {
	if x > 1 {
		return x + RecursionFunc(x-1)
	} else {
		return 1
	}
}

// 递归实现斐波那契数列
func fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(2*i) + fibonacci(i)
	}
}

/*
闭包、全局变量、局部变量：
	全局标量的特点：
		1、常驻内存
		2、可能污染全局
	局部变量的特点：
		1、不常驻内存
		2、不污染全局

	闭包：
		1、可以让一个变量常驻内存
		2、可以让一个变量不污染全局
*/

/*
闭包：
	1、闭包是指有权访问另一个函数作用域中的变量的函数
	2、创建闭包的常见方式就是在一个函数内部创建另一个函数，通过另一个函数访问这个函数的局部变量
注意：鱼鱼闭包是作用域返回的局部变量资源不会被立刻销毁回收，所以可能会占用更多的内存，过度的使用闭包会导致性能下降，建议在非常必要的时候才使用闭包。
*/

// 闭包
func getA() func() string {
	var s string = "A"
	return func() string { return s }
}

func adder2() func() int {
	var i = 10
	return func() int {
		return i
	}
}

func adder() func(y int) int {
	var i int = 10
	return func(y int) int {
		i += y
		return i
	}
}

// FuncCloser 闭包演示
func FuncCloser() {
	var i int = 22
	// 闭包
	fn := adder2()
	// 闭包
	fmt.Println(fn()) // 10
	fmt.Println(fn()) // 10
	fmt.Println(fn()) // 10
	fmt.Println(fn()) // 10
	fn2 := adder()
	fmt.Println(fn2(10)) // 20
	fmt.Println(fn2(10)) // 30
	fmt.Println(fn2(10)) // 40
	fmt.Println(fn2(10)) // 50
	fmt.Println("这里演示了闭包")
	fmt.Printf("i的值：%d\n这里看出adder和adder2中定义的i并没有对全局变量i造成什么影响\n", i)
}

// defer

/*
1、多个defer的执行顺序是“后进先出”
2、所有函数在执行return返回指令之前，都会先检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回；
3、匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值；
4、return其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）；
第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；

为啥有名返回defer语句后的修改能够修改值，而匿名返回不能返回
*/

// DeferExample 展示defer用法
func DeferExample() {
	fmt.Println("defer 会将后面跟随的语句延迟执行，多个defer语句按照FILO先进后出的顺序执行")

	fmt.Println("开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("结束")
}

// LambdaReturn 匿名返回
func LambdaReturn() {
	fmt.Println("开始")
	// 匿名自执行
	defer func() {
		fmt.Println("aaaaa")
		fmt.Println("bbbb")
	}()
	fmt.Println("结束")
}

// A 匿名返回值
func A() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

// B 有名返回值
func B() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

// C
func C() int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i, &i) // 打印结果为 c defer2: 2 0xc082008340
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i, &i) // 打印结果为 c defer1: 1 0xc082008340
	}()
	return i
}

// D 有名返回值
func D() (i int) {
	fmt.Println(&i)
	defer func() {
		i++
		fmt.Println("d defer2:", i, &i) // 打印结果为 c defer2: 2 0xc082008340
	}()
	defer func() {
		i++
		fmt.Println("d defer1:", i, &i) // 打印结果为 c defer1: 1 0xc082008340
	}()
	return i
}

// DeferFunc 有名返回值
func DeferFunc(i int) (t int) {
	defer func() {
		t += 5
	}()
	return 2
}

func FN() int {
	a := 0
	defer func(i int) {
		fmt.Println("i", i)
	}(a)
	a++
	return a
}
