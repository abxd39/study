package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	defer done()
	chan1 := make(chan chan string)
	go func(){
		chan2:=make(chan string)
		defer close(chan1)
		for {
			select {
			case  chan1 <- chan2:
				fmt.Println("这是什么鬼啊", )
			case <-ctx.Done():
				fmt.Println("退出程序1")
				return
			}
			fmt.Println("下来了！！")

			select {
			case chan2 <- "hahahha":
				fmt.Println()
			case <-ctx.Done():
				fmt.Println("退出程序2")
				return
			}
		}
	}()
	for v:= range chan1{
		if v,ok := <- v;ok{
			fmt.Println(v)
			fmt.Println("hheh")
		}
	}
	<-ctx.Done()
	fmt.Println("hello context！！")
}
