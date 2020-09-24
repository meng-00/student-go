package main

import "fmt"

// 函数
//函数存在的意义？
//函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中,给它起个名字，每次用到它的时候直接用函数名调用就可以了
//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	ret := 3
	return ret
}

//返回值可以命名也可以不命名
//命名的返回值就相当于在函数中声明的一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值可以return后省略
}

//多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

//参数的类型简写:
//当参数中多个参数的类型一致时，我么可以将非最后一个参数的类型省略
func f6(x, y int) int {
	return x + y
}

//可变长参数
//可变长参数必须放在函数参数的最后一项
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y类型是切片[]int
}

//Go语言中函数没有默认参数这个概念
func main() {
	r := sum(1, 2)
	fmt.Println(r)
	m, n := f5()
	fmt.Println(m, n)
	_, x := f5()
	fmt.Println(x)
	f7("下雨了")
	f7("下雨了", 1, 2, 3, 4, 5, 6, 7)
}
