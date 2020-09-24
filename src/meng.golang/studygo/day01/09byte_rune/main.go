package main

import "fmt"

func main() {
	s := "Hello golang"
	n := len(s)
	fmt.Println(n)
	// for _, c := range s {
	// 	fmt.Printf("%c\n", c)
	// }

	//字符修改
	s2 := "白萝卜"             //[白 萝 卜]
	s3 := []rune(s2)        //把字符串轻质转换成一个rune切片
	s3[0] = '红'             //使用单引号标识字符
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := 'H' //(int32)
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	//类型转换
	c5 := byte('H') //byte(uint8)
	fmt.Printf("c4:%T\n", c5)
	fmt.Printf("%d\n", c5)

	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
