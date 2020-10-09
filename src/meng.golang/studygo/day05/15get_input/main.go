package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入如果有空格

func useScan() {
	var s string
	fmt.Print("请输入内容:")
	fmt.Scanln(&s)
	fmt.Printf("你输的内容是:%s\n", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容:")
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是:%s\n", s)

}

func main() {
	//useScan()
	useBufio()

	// 写日志
	fmt.Fprintf(os.Stdout, "这是一条日志记录!")
	fileObj, _ := os.OpenFile("./xxx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintf(fileObj, "这是一条日志记录!")
}
