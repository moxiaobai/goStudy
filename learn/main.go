/*
Demo
*/

package main

import (
	"fmt"
	"github.com/moxiaobai/goStudy/tools"
	"sort"
	"reflect"
)

//全局变量的申明和复制
var name = "gopher"

func main() {
	fmt.Println("Hello World")
	tools.Debug()

	//函数间传递数组，数组非常大大情况下，建议用指针的方式
	arr := [5]int{1: 3, 3: 4}
	learnArray(&arr)

	learnSlice()
	learnMap()
	learnReflect()
	learnChan()
}

//学习数组
func learnArray(a *[5]int) {
	fmt.Println(a[1])

	//遍历数组
	arr1 := [...]int{1, 2, 3, 5}
	//arr1 := [4]int{1,2,3,5}
	for i, v := range arr1 {
		fmt.Printf("索引：%d, 值：%d\n", i, v)
	}
}

//学习切片，切片底层是基于数组实现
func learnSlice() {
	//长度为5，容量为10
	//slice := make([]init, 5, 10)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[2])

	for i := 0; i < len(slice); i++ {
		fmt.Printf("值:%d\n", slice[i])
	}

	newSlice := slice[1:3]
	fmt.Println(newSlice)

	newSlice1 := append(slice, 10, 11)
	fmt.Println(newSlice1)

	for _, v := range newSlice1 {
		fmt.Printf("值:%d\n", v)
	}
}

//学习Map, map基于散列来实现，可以理解为字典
func learnMap() {
	dict := make(map[string]int)
	dict["id"] = 1000
	dict["sex"] = 1
	dict["age"] = 30

	//删除键值
	//delete(dict, "age")

	//遍历是无序,也就是键值对不会按既定的数据出现
	for key, value := range dict {
		fmt.Println(key, value)
	}

	//可以按照索引排序
	var keys []string
	for key := range dict {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, dict[key])
	}
}

type User struct {
	Name string
	Age  int
}

//反射
func learnReflect() {
	u := User{"张三", 20}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Println(v)
}

//通道
func learnChan() {
	// 用make来声明一个slice，make会分配和初始化slice，map和channel。
	//无缓冲通道（同步通道）
	ch := make(chan int)

	//用go关键字开始一个并发的goroutine
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()
	fmt.Println(<-ch)

	//有缓冲通道
}