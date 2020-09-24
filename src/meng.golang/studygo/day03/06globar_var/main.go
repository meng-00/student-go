package main

import "fmt"

//变量的作用域

var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	name := "理想"
	//函数中查找变量的顺序
	//1.现在函数内部查找
	//2.找不到就往函数的外面查找,一直找到全局
	fmt.Println(x, name)
}

func main() {
	f1()
	// fmt.Println(name) 只能在函数内部使用
}
