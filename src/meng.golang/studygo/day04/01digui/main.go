package main

import "fmt"

//递归:自己调用自己
//递归适合处理那种问题相同，问题的规模越来越小的场景
//递归一定要有一个明确的推出条件
//永远不要高估自己

// 4!=4*3*2*1    =4*3!
// 5!=5*4*3*2*1	 =5*4!

//计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//上台阶的面试题
//n个台阶，一次可以走1步，也可以走2步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := f(19)
	fmt.Println(ret)
	ret2 := taijie(1024)
	fmt.Println(ret2)
}
