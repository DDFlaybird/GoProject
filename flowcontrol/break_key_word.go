package flowcontrol

import "fmt"

/*
break:
	在循环中表示跳出循环并执行后面的循环，
	switch中是执行case后的跳出语句
	合理使用break防止死循环

*/

// BreakInFor break在for用来跳出for循环
// 注意只跳出一层for循环
func BreakInFor() {
	for j := 0; j < 1000; j++ {
		for i := 0; i < 100; i++ {
			if i == 10 {
				break
			}
			fmt.Println(i)
		}
		if j == 13 {
			fmt.Println(j)
		}
	}
}

// BreakWithLable break和lable配合跳出多层循环
// 这里我们在最外层定义了一个lable位，这样在第二个for内break的时候我们直接就跳到了lable1外面。这样就可以实现break跳出多层循环(你想要跳出的循环)
func BreakWithLable() {
lableOne:
	for j := 0; j < 1000; j++ {
		for i := 0; i < 100; i++ {
			if i == 10 {
				break lableOne
			}
			fmt.Println(i)
		}
		if j == 13 {
			fmt.Println(j)
		}
	}
}

// ContinueWithfor continue 用来跳出当前循环，并执行下一循环
// 执行后可以看出打印了多次next
func ContinueWithfor() {
	for i := 0; i <= 100; i++ {
		if (i % 2) != 0 {
			fmt.Println(i)
			fmt.Println("next")
			continue
		}
	}
}

// ContinueWithLable continue和lable配合
// 和break + lable 不同。contiune + lable会在继续循环而不是直接跳出循环，
// 其实根据词义也能看出来 多余解释
func ContinueWithLable() {
lableOne:
	for i := 0; i < 3; i++ {
		for i := 0; i <= 10; i++ {
			if (i % 2) != 0 {
				fmt.Println(i)
				fmt.Println("next")
				continue lableOne
			}
		}
	}

}

// BreakInSwitch break关键字在switch中的应用
// 用来跳出switch语句；实际应用中一般情况下写不写其实都一样
func BreakInSwitch() {
	switch a := 10; a {
	case 10:
		break
	}
}

// GotoWithLable goto配合lable来使用
// 虽然可以跳到指定的lable但是执行lable后面的语句也会继续执行
// 如我们跳转到lableOne，那么lableTwo后的信息也会被打印出来
func GotoWithLable() {
	if n := 25; n > 24 {
		fmt.Println("goto开始")
		goto lableOne
	} else {
		goto lableTwo
	}
lableOne:
	fmt.Println("goto1")
lableTwo:
	fmt.Println("goto2")
}
