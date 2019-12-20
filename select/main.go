package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}

func main() {
	var str string
	switch str {
	case "A":
	case "B":
	default:
		fmt.Println("default")

	}
	ctx, _ := context.WithTimeout(context.TODO(), time.Minute)
	selectStudy(ctx)
}

func selectStudy(ctx context.Context) {
end:
	for {
		select {
		case <-ctx.Done():
			//测试break会跳到哪里 ！！！！
			//break // 只能跳出 select
			break end
		case <-time.After(time.Second):
			log.Println("轮询")
		}
		log.Println("for")
	}
	log.Println("is over!!!")
}
