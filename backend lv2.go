package main

import (
	"fmt"
)

type Movie struct {
	Name, state, introduction string
	years, time, evaluate     int
}

func main() {
	m := Movie{Name: "绿皮书"}
	n := Movie{state: "America"}
	g := Movie{years: 2018}
	h := Movie{time: 130}
	s := Movie{evaluate: 9}
	z := Movie{introduction: "讲述托尼为唐开车.唐将从纽约开始举办巡回演奏,俩人之间一段跨越种族,阶级的友谊的故事"}
	fmt.Printf("请输入你的命令\n1.获得电影名字\n2.获得地区\n3.上映时间\n4.片长\n5.评分\n6.获得简介\n7.退出程序\n")
	var option int
	for {
		fmt.Scanf("%d", &option)
		if option == 1 {
			fmt.Println(m.Name)
		} else if option == 2 {
			fmt.Println(n.state)
		} else if option == 3 {
			fmt.Println(g.years)
		} else if option == 4 {
			fmt.Println(h.time)
		} else if option == 5 {
			fmt.Println(s.evaluate)
		} else if option == 6 {
			fmt.Println(z.introduction)
		} else if option == 7 {
			return
		}
	}
}
