package main

import "fmt"

func main() {
	// 	for i := 0; i < 10; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				//设置退出标签
	// 				goto breakTag
	// 			}
	// 			fmt.Printf("%v-%v\n", i, j)
	// 		}
	// 	}
	// 	return
	// 	//标签
	// breakTag:
	// 	fmt.Println("结束for循环")

	// var flag = false
	// for i := 0; i < 10; i++ {
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break //跳出内层for循环
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)
	// 	}
	// 	if flag {
	// 		break //跳出for循环(外层的for循环)
	// 	}
	// }

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX:
	fmt.Println("over")
}
