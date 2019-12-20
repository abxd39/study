package main

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strconv"
)

func main(){

	 RandomLengthParamTest("w","a","n","g")

	//m:=map[string]interface{}{"a":"w","b":"a","c":"n"}
	//fmt.Printf("map======%q",m)
	////ctx,done:=context.WithCancel(context.Background())
	//d:=time.Now().Add(10*time.Second)
	//ctx,cancel :=context.WithDeadline(context.Background(),d)
	//obj:=WorkPool.NewDispatcher(10)
	//{
	//	obj.Run(ctx)
	//}
	//defer cancel()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 30 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}


func RandomLengthParamTest(p ...string){
	fmt.Println(reflect.TypeOf(p))
	for _,v:=range p{
		fmt.Println(v)
	}
}

func Int2float(){
	temp:=float32(12368)
	fmt.Println(reflect.TypeOf(temp))
	fmt.Println(reflect.TypeOf(1569/100))
	fl:=fmt.Sprintf("%.2f",temp/100)
	fmt.Println(fl)
	value,err:=strconv.ParseFloat(fl,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(value)
	WarningValue :=float32(temp)
	wValue:=WarningValue/100
	fmt.Println(wValue)
}