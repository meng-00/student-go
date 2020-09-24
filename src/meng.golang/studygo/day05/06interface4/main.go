package main

// 使用值接收者和指针接收者的区别?
import "fmt"

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法

// func (c cat) move() {
// 	fmt.Println("走猫步...")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("猫吃%s...\n", food)
// }

func (c *cat) move() {
	fmt.Println("走猫步...")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 animal
	c1 := cat{"tom", 4} // cat
	c2 := &cat{"加加", 4} // *cat

	// a1 = c1 // 实现animal这个接口是cat的指针类型
	a1 = &c1
	fmt.Println(a1)
	a1 = c2

	fmt.Println(a1)
}

// 使用值接收者实现接口与使用指针接收者实现接口的区别？
// 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
// 指针接收者实现接口只能存结构体指针类型的变量.
