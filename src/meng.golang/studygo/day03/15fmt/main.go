package main

import "fmt"

// fmt
func main() {
	// fmt.Print("沙河")
	// fmt.Print("娜扎")
	// fmt.Println("_____")
	// fmt.Println("沙河")
	// fmt.Println("娜扎")
	//Print("格式化字符串",值)
	//%T:查看类型
	//%d:十进制数
	//%b:二进制数
	//%o:八进制数
	//%c:字符
	//%p:指针
	//%v:值
	// var m1 = make(map[string]int, 1)
	// m1["理想"] = 100
	// fmt.Printf("%v\n", m1)
	// fmt.Printf("%#v\n", m1)
	// printBaifenbi(90)

	//字符串
	// fmt.Printf("%q\n", "李想有理想")
	// fmt.Printf("%5.2s\n", "李想有理想")

	//	获取用户输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是:", s)
	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	// fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)

}

// func printBaifenbi(num int) {
// 	fmt.Printf("%d%%\n", num)
// }
