/*
数据类型(基本类型，引用类型)

*/

package main

import "fmt"

//定义结构体
type person struct {
	age  int
	name string
}

func main() {
	name := "张三"
	fmt.Println(modifyBasicType(name))
	fmt.Println(name)

	ages := map[string]int{"张三": 20}
	fmt.Println(ages)
	modifyRefType(ages)
	fmt.Println(ages)

	//结构体可以函数基本类型传值，也可以引用类型传值
	roy := person{30, "moxiaobai"}
	fmt.Println(roy)
}

//基础类型
func modifyBasicType(s string) string {
	s = s + s
	return s
}

//引用类型，值会改变
func modifyRefType(m map[string]int) {
	m["张三"] = 10
}
