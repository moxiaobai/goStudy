package tools

import "fmt"

//对外提供的方法,函数名首字母大写，则为公开方法
func Debug() {
	print("Hello Go")
}

func Add(x, y int) int {
	z := x + y
	return z
}

//私有方法，函数名首字母小写，则为私有方法，其他包不能访问
func print(msg string) {
	fmt.Println(msg)
}
