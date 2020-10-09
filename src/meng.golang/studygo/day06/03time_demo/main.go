package main

import (
	"fmt"
	"time"
)

// 时间

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()time.Second
	ret := time.Unix(1601287323, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now+1小时
	fmt.Println(now.Add(24 * time.Hour))

	// Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02", "2020-09-30")
	if err != nil {
		fmt.Printf("paese time failed,err:%v\n", err)
		return
	}
	nextYear = nextYear.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)
	fmt.Println("------------------")
	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) // 1秒钟执行一次
	// }

	// 格式化时间 把语言中时间对象，转化成字符串类型的时间
	// 2020-09-29
	fmt.Println(now.Format("2006-01-02"))
	// 2020/09/29  14:30:20
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2020/09/29  14:30:20 AM
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	// 2020/09/29  14:30:20.111
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2020-09-29")
	if err != nil {
		fmt.Printf("paese time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep
	fmt.Println("开始sleep了")
	time.Sleep(3 * time.Second)
	fmt.Println("3秒过去了")
	n := 10
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("10秒过去了")
}

// 时区
func f2() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天这个时间
	// 按照指定格式取解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2020-10-01 11:32:00")
	// 按照东八区的时区和格式取解析一个字符串的格式
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed,err:%V\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-01 11:32:00", loc)
	if err != nil {
		fmt.Printf("paese time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	// 事件对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	// f1()
	f2()
}
