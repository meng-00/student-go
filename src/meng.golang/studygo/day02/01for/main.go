package main

import "fmt"

//流程控制之跳出for循环
func main() {
	//当i=5时就跳出for循环
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")

	//当i=5时，跳出此次for循环(不执行for循环内部的打印语句)
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
