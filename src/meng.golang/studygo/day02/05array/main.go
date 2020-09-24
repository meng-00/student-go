package main

import "fmt"

//数组
//存放元素的容器
//必须指定存放的元素类型和容量(长度)
//数组的长度是数组类型的一部分
func main() {
	// var a1 [3]bool //[true false true]
	// var a2 [4]bool //[true true false false]

	// fmt.Printf("a1:%T a2:%T\n", a1, a2)
	//数组的初始化
	//如果不初始化：默认元素为零值(布尔值:false，整型和浮点型都是0,字符串:"")
	// fmt.Println(a1, a2)
	//1.初始化方式1
	// a1 = [3]bool{true, true, true}
	// fmt.Println(a1)
	//2.初始化方式2
	// a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(a10)
	//3.初始化方式3
	// a3 := [5]int{0: 1, 4: 2}
	// fmt.Println(a3)

	//数组的遍历
	// citys := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	// for i := 0; i < len(citys); i++ {
	// 	fmt.Println(citys[i])
	// }
	//2.for range遍历
	// for i, v := range citys {
	// 	fmt.Println(i, v)
	// }

	//多维数组
	//[[1,2][3,4][5,6]]
	// var a11 [3][2]int
	// a11 = [3][2]int{
	// 	[2]int{1, 2},
	// 	[2]int{3, 4},
	// 	[2]int{4, 5},
	// }
	// fmt.Println(a11)

	//多维数组的遍历
	// for _, v1 := range a11 {
	// 	fmt.Println(v1)
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	//数组是值类型
	// b1 := [3]int{1, 2, 3} //[1 2 3]
	// b2 := b1              //[1 2 3]  Ctrl+C Ctrl+V =>把world文档从文件夹A拷贝到B
	// b2[0] = 100           //b2:[100 2 3]
	// fmt.Println(b1, b2)

	//array数组练习题
	//1.求数组[1,3,5,7,8]所有元素的和
	//找出数组中和为值指定值的两个元素的下标,比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v1 := range a1 {
		sum = sum + v1
	}
	fmt.Println(sum)

	//展出和为8的两个元素的下标为(0,3)和(1,2)
	//定义两个for循环，外层的从第一个开始遍历
	//内层的for循环从外层后面的那个开始找
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
