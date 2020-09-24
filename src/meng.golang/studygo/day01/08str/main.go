package main

import (
	"fmt"
	"strings"
)

//字符串

func main() {
	//使用/转义字符将有特殊含义的字符进行转义
	path := "\"D:\\Go\\src\\github.com\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	//多行的字符串,反引号中的数据原样输出
	s2 := `
		欲买桂花同载酒
		终不似
		少年游	
	`
	fmt.Println(s2)

	s3 := `"D:\Go\src\github.com"`
	fmt.Println(s3)
	//打印字符串长度
	fmt.Println(len(s3))
	s4 := len(s3)
	fmt.Println(s4)

	//字符串拼接
	name := "理想"
	world := "残酷"
	fmt.Println(name + world)
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	fmt.Printf("%s%s\n", name, world)

	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// //判断是否包含字符串
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))

	//前缀和后缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s5 := "abcdeb"
	fmt.Println(strings.Index(s5, "c")) //c在字符串中在第二位
	fmt.Println(strings.LastIndex(s5, "b"))

	//join字符串拼接
	fmt.Println(strings.Join(ret, "+"))
}
