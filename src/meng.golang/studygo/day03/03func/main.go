package main

import "fmt"

//函数:一段代码的封装
func f1() {
	fmt.Println("Hello 沙河!")
}

func f2(name string) {
	fmt.Println("Hello,", name)
}

//带参数和返回值的函数
func f3(x int, y int) int {
	sum := x + y
	return sum
}

//参数类型简写
func f4(x, y int) int {
	return x + y
}

//可变参数
func f5(title string, y ...int) int {
	fmt.Println(y)
	return 1
}

//命名返回值
func f6(x, y int) (sum int) {
	sum = x + y //如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      //return后面可以省略返回值变量
}

//Go语言支持多个返回值
func f7(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("理想")
	f2("籍无名")
	f3(12, 30) //调用函数
	fmt.Println(f3(32, 543))
	ret := f3(3412, 5643)
	fmt.Println(ret)
	fmt.Println(f6(5, 8))
	f5("lixiang", 1, 2, 3, 4, 5, 6, 7)
}
