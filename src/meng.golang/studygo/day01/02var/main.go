package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明

var (
	name string
	age  int
	isOK bool
)

func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	name = "golang"
	age = 15
	isOK = true
	//go语言中非全局变量声明必须使用，不使用就编译不过去
	fmt.Print(isOK)               //在终端中输出要打印的内容
	fmt.Printf("name:%s\n", name) //%s：占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)              //打印完指定的内容之后会在后面加一个换行符

	//声明变量同时赋值
	var s1 string = "let's go"
	fmt.Print(s1)

	//类型推导
	var name = "xiaoming"
	var num = 10
	fmt.Printf(name)
	fmt.Print(num)

	// s1 := "haha"   同一个作用域中不能重复声明同名的变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
