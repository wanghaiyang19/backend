package main

import (
	"fmt"
)

func Xingqiji() chan string {
	ch := make(chan string, 2)

	go func() {
		i := "今天是疯狂星期四"
		ch <- i

		close(ch)
	}()

	return ch
}

func Paopao(ch chan string) string {
	sum := "请vivo 50"
	for v := range ch {
		sum = v + "," + sum
	}
	return sum
}

func main() {
	ch := Xingqiji()
	res := Paopao(ch)
	fmt.Print("泡泡：")
	fmt.Println(res) // 25

}
