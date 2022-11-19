package main

import (
	"context"
	"fmt"
	"github.com/eiannone/keyboard"
	"runtime"
	"time"
)

var (
	Zhuangdan = make(chan struct{}, 1)
	Miaozhun  = make(chan struct{}, 1)
	Shot      = make(chan struct{}, 1)
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fire(ctx)
	var (
		input rune
		err   error
	)
	for input != 'q' {
		input, _, err = keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("停止前 Goroutine 数量：%d\n", runtime.NumGoroutine())
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("打炮结束.")
	fmt.Printf("停止后 Goroutine 数量：%d\n", runtime.NumGoroutine())
}

func fire(ctx context.Context) {
	for i := 0; i < bulletNum; i++ {
		go func(pos int) {
			for {
				select {
				case <-Zhuangdan:
					time.Sleep(time.Second)
					fmt.Print("装弹 ->")
					Miaozhun <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("装弹手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	for i := 0; i < aimNum; i++ {
		go func(pos int) {
			for {
				select {
				case <-Miaozhun:
					time.Sleep(time.Second)
					fmt.Print(" 瞄准 ->")
					Shot <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("瞄准手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	for i := 0; i < shotNum; i++ {
		go func(pos int) {
			for {
				select {
				case <-Shot:
					time.Sleep(time.Second)
					fmt.Println(" 发射！")
					Zhuangdan <- struct{}{}
				case <-ctx.Done():
					fmt.Printf("发射手%d：回家了\n", pos)
					return
				}
			}
		}(i)
	}
	Zhuangdan <- struct{}{}
}
