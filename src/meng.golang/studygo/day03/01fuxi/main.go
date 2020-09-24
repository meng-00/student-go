package main

import "fmt"

//day03f复习

func main() {
	// var name string
	// name = "理想"
	// fmt.Println(name)
	// var ages [30]int //声明了一个变量ages,它是[30]int类型
	// ages = [30]int{1, 2, 3, 4, 5}
	// fmt.Println(ages)
	// //第二种类型初始化
	// var ages2 = [...]int{1, 2, 3, 4}
	// fmt.Println(ages2)
	// var ages3 = [...]int{1: 2, 3: 4}
	// fmt.Println(ages3)

	//二维数组注意事项
	// var a1 = [...][2]int{
	// 	[2]int{2, 3},
	// 	[2]int{5, 6},
	// 	[2]int{9, 31},
	// }
	// fmt.Println(a1)
	//多维数组只有最外层可以使用...

	//数组是值类型
	// x := [3]int{1, 2, 3}
	// fmt.Println(x) //[1,2,3]
	// f1(x)
	// fmt.Println(x) //?[1,2,3]

	//切片(Slice)：
	// var s1 []int //没有分配内存,== nil
	// fmt.Println(s1)
	// fmt.Println(s1 == nil)
	// s1 = []int{1, 2, 3}
	// fmt.Println(s1)

	//make初始化 分配内存
	// s2 := make([]bool, 2, 5) //2容量，5长度
	// fmt.Println(s2)          //[false,false]
	// s3 := make([]int, 0, 4)
	// fmt.Println(s3 == nil) //此数值为false，因为make已经分配了内存

	// s1 := []int{1, 2, 3} //[1 2 3]
	// s2 := s1
	// // var s3 []int //此时s3为nil，未对其初始化
	// var s3 = make([]int, 0, 3) //此切片容量为0，不能赋值数据到此切片
	// copy(s3, s1)               //copy也布恩那个对其切片进行扩容
	// fmt.Println(s2)            //[1 2 3]
	// s2[1] = 200
	// fmt.Println(s2) //[1 200 3]
	// fmt.Println(s1) //[1 200 3]
	// fmt.Println(s3) //[]

	// var s1 []int
	// // s1[0] = 100   //没有对切片进行初始化，所以不能进行直接对其赋值
	// // fmt.Println(s1)
	// s1 = append(s1, 1)   //append函数会自动对切片自动进行初始化
	// fmt.Println(s1)

	//指针
	//Go语言中的函数传递的都是值(Ctrl+c Ctrl+v)
	// addr := "沙河"
	// addrP := &addr
	// fmt.Println(addrP) //内存地址
	// fmt.Printf("%T\n", addrP)
	// addrV := *addrP //根据内粗地址找值
	// fmt.Println(addrV)

	//map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10) //对map进行初始化，如果存储的值超过了申请的内存空间，则会自动扩容
	m1["lixiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) //如何key不存在返回的是value对应类型的零值

	//如果返回的值是布尔型,我们通常用ok去接受它
	score, ok := m1["lixiang"]
	if !ok {
		fmt.Println("没有这个人")
	} else {
		fmt.Println("分数是", score)
	}
	delete(m1, "xiaoming") //删除的key不存在什么都不干
	delete(m1, "lixiang")
	fmt.Println(m1)
	fmt.Println(m1 == nil) //已经被初始化，开辟了内存空间,所以不等于nil
}

func f1(a [3]int) {
	//GO语言中的函数传递的都是值(Ctrl+c Ctrl+v)
	a[1] = 100
}
