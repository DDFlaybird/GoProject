package datastructure

import (
	"fmt"
	"reflect"
	"sort"
)

/*
创建map类型：
	1、make创建map类型
		mapA := make(map[string]string)
		mapA["username"] = "zhangsan"
		mapA["age"] = "22"
	2、推导式初始化
		mapA := map[string]string{"username": "zhangsan", "age": "22"}
	3、直接声明
		var mapA = map[string]string{"username": "zhangsan", "age": "22"}
*/
var Mapa = map[string]string{
	"username": "zhangsan",
	"age":      "22",
}

// ForRangeInMap 使用range遍历map中的数据
func ForRangeInMap() {
	mapA := map[string]string{"username": "张三", "age": "22"}
	for k, v := range mapA {
		fmt.Printf("K:%v, v:%v\n", k, v)
	}
	// 你还可以像python中的dict一样获取数据
	fmt.Printf("mapA['username']: %v\n", mapA["username"])
}

// CURDInMap Map中的增删改查
func CURDInMap() {
	var userinfo = map[string]string{
		"username": "张三",
		"city":     "guangdong",
	}
	userinfo["sex"] = "男"
	fmt.Println("新增和更新")
	fmt.Printf("map对象的新增：userinfo[‘sex’]='男',\n增加后userinfo的值为：%v\n", userinfo)
	userinfo["username"] = "李四"
	fmt.Printf("map对象的更新：userinfo[‘username’]='男',\n增加后userinfo的值为：%v\n", userinfo)
	fmt.Println("获取")
	fmt.Printf("可以直接像python中的dict一样使用下角标来获取map中的信息,\nuserinfo['username']: %v\n", userinfo["username"])
	v, err := userinfo["username"]
	fmt.Printf("或者使用 v, err := userinfo['username']，\n其中v:%v是值，err:%v 是错误信息\n", v, err)
	fmt.Println("如果有值那么err就返回true，如果值不存在就返回false")
	fmt.Println("删除")
	fmt.Println("使用delete()删除map中的键值对")
	delete(userinfo, "sex")
	fmt.Printf("使用delete()删除sex: delete(userinfo, 'sex'),\nuserinfo指的是map对象，“sex”指需要删除键值对的key，删除后的值为：%v\n", userinfo)
}

// IntoMap 详解map
// 三种声明map的方式中只有userinfo中指定了长度和容量的map对象才有初始值(nil)
// 不然仅仅是个map对象，这是因为make(map[string]string，3， 3)生成了一个内容为map的列表
// 注意 userinfo这个对象虽然打印出来列表内部的元素显示的是map[] 但是实际上你不能给这个map对象
// 进行增删改查，所以当我们想要给userinfo内部的map添加元素的时候我们需要再次声明一下map对象
func IntoMap() {
	var user map[string]string
	info := make(map[string]string)
	userinfo := make([]map[string]string, 3)
	fmt.Println(user)
	fmt.Println(info)
	fmt.Println(userinfo)
	fmt.Println()
	for _, v := range userinfo {
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(v == nil)
	}
	userinfo[0] = make(map[string]string)
	userinfo[0]["username"] = "zhangsan"
	fmt.Println(userinfo)
	// map对象是引用类型
	userinfo2 := userinfo
	userinfo2[0]["username"] = "李四"
	// userinfo2是userinfo的副本
	// 改变userinfo2中的username对应的userinfo中的username也会改变
	fmt.Println(userinfo2)
	fmt.Println(userinfo)
	// map对象的排序
	// mapsort也是一个map对象，这个时候size的意思表示的这个map分配的内存空间。
	mapsort := make(map[int]int, 10)
	mapsort[11] = 101
	mapsort[10] = 100
	mapsort[5] = 99
	mapsort[3] = 76
	mapsort[1] = 35
	fmt.Println(mapsort)
	// 这里建议sort包
	// 1、使用map键值对的特点将键进行排序并保存到一个切片中，随后遍历切片再map中取值
	var keyslice []int
	for k, _ := range mapsort {
		keyslice = append(keyslice, k)
	}
	fmt.Println(keyslice)
	// sort排序 但是需要先将slice定义为sort包中的intslice类的
	sort.Sort(sort.IntSlice(keyslice))
	fmt.Println(keyslice)
	// 稳定排序
	sort.Stable(sort.Reverse(sort.IntSlice(keyslice)))
	fmt.Println(keyslice)
	// 这里我们可以直接用Ints, 就不需要再转化一下类型了。
	sort.Ints(keyslice)
	fmt.Println(keyslice)
	for _, v := range keyslice {
		fmt.Printf("Key:%d, value:%d\n", v, mapsort[v])
	}
}
