package main

import (
	"context"
	"fmt"
	"time"
)

var battle = make(chan string)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	ch := make(chan int)
	go func() {
		produce(ctx, ch)
	}()
	go func() {
		consume(ch)
	}()

	<-ctx.Done()
	time.Sleep(time.Second)

}

func produce(ctx context.Context, pch chan<- int) {
	//如果试图从 pch 读取数据 编译不过的
	//date:=<-pch
	i:=0
	for {
		select {
		case <-ctx.Done():
			close(pch)
			return
		default:
			time.Sleep(time.Second)
		}
		pch <- i
	}

}

func consume(cch <-chan int) {
	//如果试图从 pch 读取数据 编译不过的
	//cch<- struct{}{}
	for {
		select {
		case date, ok := <-cch:
			if !ok {
				return
			}
		 _=date
		}
	}
}

//for true {
//	select {
//	case <-tem:
//		fmt.Println("struct")
//	case <-time.After(time.Second * 5):
//		fmt.Println("5")
//	case <-time.After(time.Second * 10):
//		fmt.Println("10")
//		return
//	}
//}
//fmt.Println("main over")
//
//done := make(chan struct{})
//langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
//for _, l := range langs {
//	go warrior(l, done)
//}
//for _ = range langs {
//	<-done
//}

func producer(ctx context.Context) chan chan int {
	ints := make(chan chan int)
	go func() {
		ch := make(chan int, 100)
		i := 1
		defer func() {
			fmt.Println("producer 退出")
			ch <- 888
			close(ch)
		}()
		defer func() {
			fmt.Println("关闭通道ints")
			close(ints)
		}()
		timeOut := time.After(time.Second * 5)
		for true {
			select {
			case <-ctx.Done():
			case ints <- ch:
				fmt.Println("压入")
			default:
				fmt.Printf("a %v\n", i)

			}
			ch <- i
			i++
			select {
			case <-ctx.Done():
			case t := <-timeOut:
				fmt.Printf("过期时间到了 %v\n", t.Format("2006-01-02 15:04:05"))
				return
			default:
				fmt.Printf("default %v\n", i-1)
			}
			time.Sleep(time.Second * 3)
		}
	}()
	return ints
}

func consume1(ints chan chan int) {
	for in := range ints {
		for true {
			i, Ok := <-in
			if !Ok {
				fmt.Println("channel closed!")
				return
			}
			fmt.Printf("------->%v\n", i)
		}

	}
}

func warrior(name string, done chan struct{}) {
	select {
	case opponent := <-battle:
		fmt.Printf("%s beat %s\n", name, opponent)
	case battle <- name:
		// I lost :-(
		//fmt.Println(name)
	}
	done <- struct{}{}
}
