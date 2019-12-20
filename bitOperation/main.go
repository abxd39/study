package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

func main()  {
	var count  int8
	count =127
	fmt.Printf("%b\n",count)
	temp:=30
	fmt.Printf("%b\n",temp)
	fmt.Printf("%b\n", temp & 0xF)
	max,err:=strconv.ParseInt("0x7FFFFFFFFFFFFFFF",0,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(max)
	maxU,err:=strconv.ParseUint("0x7FFFFFFFFFFFFFFF",0,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(maxU)
	var i int
	var ui uint
	var i64 int64
	var u64 uint64
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(ui))
	fmt.Println(unsafe.Sizeof(i64))
	fmt.Println(unsafe.Sizeof(u64))
	strTimestamp :=fmt.Sprintf("%v",time.Now().Unix())
	fmt.Println(strTimestamp)
	timestamp, err := strconv.ParseInt(strTimestamp, 10, 64)
	first,second:=err.(*strconv.NumError)
	fmt.Printf("%+v\n",second)
	if second{
		if first.Err!=nil{
			fmt.Println(err)
		}
		fmt.Println("01",first.Func)
		fmt.Println("02",first.Num)
	}

	fmt.Println(timestamp)

	maxUint64:= strconv.FormatUint(math.MaxUint64,10)
	fmt.Println(maxUint64)
	//maxUint64Itoa:=strconv.Itoa(int(math.MaxUint64))
	fmt.Println(uint(math.MaxUint64))
	}
