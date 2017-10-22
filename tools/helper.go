package tools

import "fmt"

//对外提供的方法
func Debug() {
	print("Hello Go")
}

func Add(x, y int) int {
	z := x + y
	return z
}

//私有方法
func print(msg string) {
	fmt.Println(msg)
}
