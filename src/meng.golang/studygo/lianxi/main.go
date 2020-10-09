package main

import "fmt"

type Phone interface {
	call(name string)
}
type Person struct {
	Name string
}

//方法
func (p *Person) Test2(newName string) string {
	p.Name = newName
	return p.Name
}

//接口
func (p Person) call(name string) string {
	return p.Name + "打电话给" + name
}

func main() {
	p1 := Person{"abc"}
	str := p1.call("xyz")
	fmt.Println(str) //abc打电话给xyz
}
