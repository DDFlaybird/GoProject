package FlowControl

import "fmt"

// 最简单的if-else语句
func SimpleIfExample(flag bool) bool {
	if flag {
		return true
	} else {
		return false
	}
}

// 直接再if语句中声明一个新的变量并进行比较
// for 语句中也有类似的，利用的是同一行代码的执行顺序是从左到右
// 这里需要注意一点在if语句内部声明的变量是一个局部变量再if外部不能访问到
func AnthorSimpleIfExample(ages int) {
	if age := 18; age > ages {
		fmt.Println("未成年")
		fmt.Println("可以调用if语句中声明的变量age")
		fmt.Println("age:", age)
	} else {
		fmt.Println("成年")
		fmt.Println("可以调用if语句中声明的变量age")
		fmt.Println("age:", age)
	}
}

// elif的写法
func ElifSimpleExample(ages int) {
	if age := 18; age > ages {
		fmt.Println("未成年")
	} else if middleAge := 30; middleAge > ages {
		fmt.Println("青年")
	} else if oldAge := 50; oldAge > ages {
		fmt.Println("中年")
	} else {
		fmt.Println("老年")
	}
}
