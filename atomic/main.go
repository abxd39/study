package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var count int64

func main(){
	//w:=sync.WaitGroup{}
	//w.Add(100)
	//atomic.LoadInt64(&count)
	//if count==0{
	//	count =6100001
	//	fmt.Println("~~~~~~~~~~~")
	//}
	//for i:=0;i<100;i++{
	//	go func(int2 int) {
	//		atomic.AddInt64(&count,1)
	//		atomic.StoreInt64(&count,count)
	//		couponCountStr:=fmt.Sprintf("%d",count)
	//		year:=time.Now().Format("2006010215")
	//		year="31"+year[2:]+couponCountStr+"1"
	//		fmt.Printf("code=%v,gorutione=%v\r\n",year,int2)
	//		w.Done()
	//	}(i)
	//
	//}
	//w.Wait()
	CouponCount,err:=strconv.ParseUint("600000001",10,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(CouponCount)
	str:=",1,12,13,15,19,,,16,,"
	tagList := strings.Split(str, ",")
	tagm := make(map[string]int, 0)
	for _, v := range tagList {
		if v==""{
			continue
		}
		tagm[v]++
	}
	fmt.Printf("%+v\r\n",tagm)
	strNew := ""
	for key, _ := range tagm {
		strNew += fmt.Sprintf("%v,", key)
	}

	fmt.Printf("%v\r\n",strNew)
	if strNew!=""{
		fist:=strNew[:1]
		if fist==","{
			str=strNew[1:]
		}
		second:=strNew[len(strNew)-1:]
		if second==","{
			strNew =strNew[:len(strNew)-1]
		}
	}
	fmt.Printf("%v\r\n",strNew)
	time.Sleep(time.Second*10)
}
