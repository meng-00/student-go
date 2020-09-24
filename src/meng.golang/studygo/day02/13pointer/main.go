package main

import "fmt"

//pointer

func main() {
	//1.&:取地址
	//2.*:根据地址取值
	// n := 18
	// p := &n
	// fmt.Println(p)
	// fmt.Printf("%T\n", p) //*int:int类型的指针
	// //2.*:根据地址取值
	// m := *p
	// fmt.Println(m)
	// fmt.Printf("%T\n", m)

	// var a *int //nil pointer
	// *a = 100  //此时赋值会错误
	// fmt.Println(*a)

	//new函数申请一个内存地址
	var a2 = new(int)
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)

}

//make和new的区别
//1.make和nwe都是用来申请内存的
//2.new很少用，一般用来给基本数据类型申请内存，string、int,返回的是对应类型的指针(*string、*int)
//3.make是用来给slice、map、chan申请内存的，make函数返回的是对应的这三个类型本身
