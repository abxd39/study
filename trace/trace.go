package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

//https://www.sohu.com/a/326924107_657921?spm=smpc.author.fd-d.1.1563519608307sMDmEU1
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	fmt.Println(`
	View trace：查看跟踪 \
	Goroutine analysis：Goroutine 分析
	Network blocking profile：网络阻塞概况
	Synchronization blocking profile：同步阻塞概况
	Syscall blocking profile：系统调用阻塞概况
	Scheduler latency profile：调度延迟概况
	User defined tasks：用户自定义任务
	User defined regions：用户自定义区域
	Minimum mutator utilization：最低 Mutator 利用率
	`)
	ch := make(chan string)
	go func() {
		ch <- "wyw"
	}()
	<-ch
}
