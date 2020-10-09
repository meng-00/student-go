package main

import (
	"time"

	"meng.golang/studygo/day06/mylogger"
)

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("debug")
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}
