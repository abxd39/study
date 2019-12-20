package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
	//numThread:=debug.SetMaxThreads(3)
	//fmt.Printf("内核线程为%v",numThread)
	n := runtime.GOMAXPROCS(2)
	fmt.Printf("gorountine 上线为 %v个", n)
}

type newWork struct {
	Name string
}

var OrderNumber int32

func (p *newWork) Run() {
	atomic.AddInt32(&OrderNumber, 1)
	fmt.Printf("%v %v\n", p.Name, OrderNumber)
}

const MaxThread = 100

func main() {
	go func() {
		http.ListenAndServe(":6060",nil)
	}()
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//goroutineStudy(ctx)
	//ContextStudy(ctx)
	go func() {
		i := 0
		for true {
			select {
			case <-ctx.Done():
			default:
			}
			i++
			_ = i
		}
	}()
	go func() {
		i := 0
		for true {
			select {
			case <-ctx.Done():
			default:
			}
			i++
			_ = i
		}
	}()
	go func() {
		i := 0
		for true {
			select {
			case <-ctx.Done():
			default:
			}
			i++
			_ = i
		}
	}()
	go func() {
		i := 0
		for true {
			select {
			case <-ctx.Done():
			default:
			}
			i++
			_ = i
		}
	}()
	<-ctx.Done()
	time.Sleep(time.Second * 70)

	ch:=make(chan string,2)

	f1:=func() {
		ch<- "正在处理"
		defer close(ch)
	}

	go f1()
	for v:=range ch{
		fmt.Printf("知道了，我%v",v)
	}

	fmt.Println("sleep")
	time.Sleep(time.Minute)
}

func goroutineStudy(ctx context.Context) {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//select 只会执行一次 如果灭有default 则会阻塞进程
		select {
		case <-ctx.Done():
			fmt.Println("goroutine 成功退出^V^!!")
			time.Sleep(time.Second * 2)
		}
		fmt.Println("select 退出了")
		w.Done()
	}()

	w.Wait()
	defer func() {
		fmt.Println("main 安全退出")
	}()
}

func ContextStudy(ctx context.Context) {
	gen := func(ctx1 context.Context) <-chan int {
		dst := make(chan int)

		go func() {
			n := 1
			fp, err := os.OpenFile("goroutine.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				close(dst)
				return
			}
			defer func() {
				fp.Close()
				log.Println("文件已经关闭")
			}()
			for {
				log.Println(n)
				str := "0"
				select {
				case <-ctx1.Done():
					log.Println("ctx 时间到！！！")
					return
				case dst <- n:
					n++
					time.Sleep(time.Second)

					if n > 6 {
						log.Println(n)
						str = fmt.Sprintf("go n -->%v\r\n", <-dst)
						log.Println(n)
					} else {
						str = fmt.Sprintf("go n -->%v\r\n", n)
					}

				case <-time.After(time.Second):
					n++
					log.Println("定时器")
					str = fmt.Sprintf("go 定时器 -->%v\r\n", n)
					//_,ok :=<-dst
				case _, ok := <-dst:
					if !ok {
						log.Println("dst 已经关闭")
					} else {
						log.Println("dst 正常！！")
					}
				}

				fp.WriteString(str)

			}
		}()
		return dst
	}
	for n := range gen(ctx) {
		log.Println(n)
		if n == 5 {
			break
		}
	}
}

func TryE() {
	errs := recover()
	if errs == nil {
		fmt.Println("没有捕获到Ctr+c 消息")
		return
	}
	exeName := os.Args[0]                                             //获取程序名称
	now := time.Now()                                                 //获取当前时间
	pid := os.Getpid()                                                //获取进程ID
	time_str := now.Format("20060102150405")                          //设定时间格式
	fname := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, time_str) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
	fmt.Println("dump to file ", fname)
	f, err := os.Create(fname)
	if err != nil {
		fmt.Println(errs)
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%v\r\n", errs)) //输出panic信息
	f.WriteString("========\r\n")
	f.WriteString(string(debug.Stack())) //输出堆栈信息
}


//iv= y3BtTsIiaU1txfme
//auth_code= B6swsxAzX1e0hIq9i0WD3TpeLNhUgOM0uHCx1znTVijX91j4VRTFrD178Q9j9%2BRX
