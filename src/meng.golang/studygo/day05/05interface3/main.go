// type 接口名 interface {
// 	方法名1(参数1,参数2...)(返回值1,返回值2...)
//  方法名1(参数1,参数2...)(返回值1,返回值2...)
// ...
// }

// 用来给变量\参数\返回值设置类型
// 接口的实现
// 一个变量如果实现了接口中规定的所有的方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量

package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步...")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	name string
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡冻")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s!\n", food)
}

func main() { // 定义一个接口类型的变量
	var a1 animal
	bc := cat{ //定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}

	kfc := chicken{
		name: "淘气",
		feet: 2,
	}
	a1 = bc
	a1.eat("小黄花")
	fmt.Printf("%T\n", a1)
	a1.move()
	fmt.Println(a1)

	a1 = kfc
	fmt.Printf("%T\n", a1)
}

// 接口保存的分为两部分:
// 值的类型和值本身
