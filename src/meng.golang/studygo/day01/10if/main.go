package main

import "fmt"

//if条件判断
func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家显示功能赌场开业了!")
	// } else {
	// 	fmt.Println("改写暑假作业了")
	// }

	// //多个判断条件
	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("好好学习")
	// }

	//作用域
	//age变量此时只在if条件判断语句中生效
	if age := 10; age > 18 {
		fmt.Println("澳门首家显示功能赌场开业了!")
	} else {
		fmt.Println("改写暑假作业了")
	}
}
