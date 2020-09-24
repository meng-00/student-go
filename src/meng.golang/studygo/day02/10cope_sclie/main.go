package main

import "fmt"

func main() {
	// a1 := []int{1, 3, 5}
	// a2 := a1
	// // var a3 []int //nuil
	// var a3 = make([]int, 3, 3)
	// copy(a3, a1) //copy
	// fmt.Println(a1, a2, a3)
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// //将a1中的索引为1的3这个元素删掉
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1)
	// fmt.Println(cap(a1)) //切片不存值，底层数组才存值

	x1 := [...]int{1, 3, 5} //数组
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1)
}
