package main

import (
	"fmt"

	"meng.golang/studygo/day05/calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行!")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
