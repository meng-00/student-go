package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(len(s))

	// for i, v := range s {
	// 	fmt.Printf("%d %c\n", i, v)
	// }
	// for _, v := range s {
	// 	fmt.Printf("%c\n", v)
	// }

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		for j := 10; j >= 9; j-- {
			fmt.Printf("%d*%d=%d\t", i, j, j*i)
		}
		fmt.Println()
	}
}
