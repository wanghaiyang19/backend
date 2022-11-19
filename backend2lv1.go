package main

import (
	"fmt"
	"time"
)

func main() {

	go callback(func(好康的 string) {
		fmt.Println(好康的)
	})
	打电动()
	time.Sleep(6 * time.Second)
}

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() {
	fmt.Println("输了啦，都是你害的")
}

func callback(callback func(string)) {
	好康的 := 欢迎来我家玩()
	callback(好康的)
}
