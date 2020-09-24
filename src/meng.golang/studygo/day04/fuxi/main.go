package main

import "fmt"

func yuanshuai(name string) {
	fmt.Println("Hello", name)
}

func lixiang(f func(string), name string) {
	f(name)
}

func zhoulin() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func low(f func()) {
	f()
}

//闭包
func b1(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main() {
	lixiang(yuanshuai, "理想")
	ret := zhoulin()
	fmt.Printf("%T\n", ret)
	sum := ret(10, 20)
	fmt.Println(sum)

	//闭包示例
	fc := b1(yuanshuai, "元帅")
	low(fc)
}
