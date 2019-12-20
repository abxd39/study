package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	for i:=0;i<100;i++{
		//生成长度为8的随机数字
		randstr:=fmt.Sprintf("%010v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
		fmt.Println(randstr)
		time.Sleep(time.Nanosecond*10)

	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i :=0;i<10;i++ {
		fmt.Printf("other %010v\r\n",rd.Intn(300))
	}
}

