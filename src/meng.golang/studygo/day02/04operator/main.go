package main

import "fmt"

//运算符
func main() {
	// var (
	// 	a = 2
	// 	b = 3
	// )
	//算术运算符
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)
	// a++ //单独的语句,不能放在=的右边赋值   ==>a=a+1
	// b-- //单独的语句,不能放在=的右边赋值   ==>b=b-1
	// fmt.Println(a)

	//关系运算符
	// fmt.Println(a == b)
	// fmt.Println(a != b)
	// fmt.Println(a >= b)
	// fmt.Println(a > b)
	// fmt.Println(a <= b)
	// fmt.Println(a < b)

	//逻辑运算符
	//如果年龄大于18岁并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班的!")
	} else {
		fmt.Println("不用上班!")
	}
	//如果年龄小于18岁或者年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("苦逼上班的!")
	} else {
		fmt.Println("不用上班!")
	}

	//not取反,原来为真就为假，原来为假就为真
	isMarried := false
	fmt.Println(isMarried)  //false
	fmt.Println(!isMarried) //true

	// 位运算
	// 5的二进制标识:101
	// 2的二进制表示:10

	// &:按位与
	fmt.Println(5 & 2)
	// |:按位或
	fmt.Println(5 | 2)
	// >>:将二进制位右移指定位数
	fmt.Println(5 << 1)

	var m = int16(1)
	fmt.Println(m << 10)

	var x int
	x = 10
	x += 1 //可直接写成x++
	x -= 1 //可直接写成x--
	x *= 2
	x /= 2
	x %= 2

	x <<= 2 // x = x<<2
	x &= 2  // x = x&2
	x |= 3  // x = x|3
	x >>= 3 // x = x>>2

}
