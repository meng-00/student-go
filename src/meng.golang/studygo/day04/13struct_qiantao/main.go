package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

// type person struct {
// 	name string
// 	age  int
// 	addr address
// }

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体
	// address:address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周琳",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	// fmt.Println(p1.city) // 现在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找
	fmt.Println(p1.address.city)
}
