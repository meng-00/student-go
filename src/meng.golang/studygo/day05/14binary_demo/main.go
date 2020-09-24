package main

import "fmt"

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

// 二进制实用途
// 111
// 左边1表示吃饭
// 中间的1表示睡觉
// 右边的1便是打豆豆
func f(agr int) {
	fmt.Printf("%b\n", agr)
}

func main() {
	f(eat | da)
	f(eat | da | sleep)
}
