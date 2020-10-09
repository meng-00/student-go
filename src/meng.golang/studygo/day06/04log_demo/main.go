package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log demo

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err%d\n", err)
		return
	}
	log.SetOutput(fileObj)
	// log.SetOutput(os.Stdout)    // 将日志信息输出都终端
	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 3)
	}
}
