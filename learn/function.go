/*
函数方法
函数是指不属于任何结构体、类型的方法，函数是没有接收者的
方法是有接收者
*/

package main

import "fmt"

type animal struct {
	name string
}

func main() {
	p := animal{name: "小八"}
	(&p).modify()
	fmt.Println(p.ShowName())

	sum, mount := sum(1, 2)
	fmt.Println(sum, mount)

	print("1", "2", "3")
}

//私有函数
func add(a, b int) int {
	return a + b
}

//方法,有接收者 （值接收者和指针接收者）
func (p animal) ShowName() string {
	return "The Animal name is " + p.name
}

func (p *animal) modify() {
	p.name = "小五"
}

//多值返回
func sum(a, b int) (int, int) {
	return a + b, a * b
}

//可变参数
func print(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}
