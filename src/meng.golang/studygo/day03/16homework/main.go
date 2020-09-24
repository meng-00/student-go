package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下:
a.名字中每包含1个'e'或'E'分1枚金币
b.名字中每包含1个'i'或'I'分2枚金币
c.名字中每包含1个'o'或'O'分3枚金币
d.名字中每包含1个'u'或'U'分4枚金币
*/

var (
	coins = 5000
	users = []string{
		"Mattew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	destribution = make(map[string]int, len(users))
)

func dispathCion() (left int) {
	//1.依次那个每个人的名字
	for _, name := range users {
		//2.拿个一个人名根据分金币的规则去分金币,
		for _, c := range name {
			switch c {
			case 'e', 'E':
				//满足这个条件分1个金币
				destribution[name]++
				coins--
			case 'i', 'I':
				//满足这个条件分2个金币
				destribution[name] += 2
				coins -= 2
			case 'o', 'O':
				//满足这个条件分3个金币
				destribution[name] += 3
				coins -= 3
			case 'u', 'U':
				//满足这个条件分4个金币
				destribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

func main() {
	left := dispathCion()
	fmt.Println("剩下", left)
	for k, v := range destribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}
