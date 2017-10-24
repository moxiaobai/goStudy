/*
接口
*/

package main

import "fmt"

//接口的声明
type animals interface {
	printInfo()
}

// 定义Stringer为一个接口类型，有一个方法String
type Stringer interface {
	String() string
}

// 定义pair为一个结构体，有x和y两个int型变量。
type pair struct {
	x, y int
}

func main() {
	var a animals
	var c cat
	invoke(c)

	a = c
	a.printInfo()

	var d dog
	a = d
	a.printInfo()

	// 调用pair类型p的String方法
	p := pair{3,4}
	fmt.Println(p.String())

	// 声明i为Stringer接口类型
	var i Stringer
	i = p
	fmt.Println(i.String())
}

// 定义pair类型的方法，实现Stringer接口。
func (p pair) String() string {  // p被叫做“接收器”
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func invoke(a animals) {
	a.printInfo()
}

//一般类型的声明
type cat int
type dog int

func (c cat) printInfo() {
	fmt.Println("a cat")
}

func (c dog) printInfo() {
	fmt.Println("a dog")
}
