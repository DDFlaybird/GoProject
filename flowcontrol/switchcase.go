package flowcontrol

import "fmt"

// SwitchExample 简单的switch演示
func SwitchExample(grade string) {
	switch grade {
	case "A":
		fmt.Println("优秀")
		break // go中break可以写也可以不写
	case "B", "c":
		// switch 中可以和多个参数对比 符合其中一个就进入这个case
		fmt.Println("合格")
		break
	case "D":
		fmt.Println("不合格")
		break
	default: // default是当输出的都不符合的时候就进入这一层执行default中的内容
		fmt.Println("太垃圾了没有分类")
		break
	}
}

// SwitchCaseExampleB 直接在switch中声明变量
func SwitchCaseExampleB() {
	switch a := 20; a {
	case 20, 30, 40:
		fmt.Println("20, 30 , 40之中")
		fallthrough // 穿透 有这个语句那么这个case执行完后 后面跟着的case也会执行，但是一个fallthrough只能穿透一层 如果想要穿透多层则在对应的case中添加fallthrough关键字
	case 10:
		fmt.Println("10中")
	}
}

// SwitchCaseExampleC 在Switch中使用表达式
// 如果用表达式那么不建议用switch用if来实现
func SwitchCaseExampleC(age int) {
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case 24 < age || age <= 60:
		fmt.Println("好好搬砖")
	}
}
