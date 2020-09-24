package main

import "fmt"

func main() {
	//十进制
	var i1 = 10177
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制数转换成二进制数
	fmt.Printf("%o\n", i1) //把十进制数转换成八进制数
	fmt.Printf("%x\n", i1) //把十进制转换成十六进制
	//八进制
	i2 := 075
	fmt.Printf("%d\n", i2)
	i3 := 0xef1234
	fmt.Printf("%d\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}
