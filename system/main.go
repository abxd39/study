package main

import (
	"fmt"
	"runtime"
)

func main (){
	nc:=runtime.NumCPU()
	fmt.Println(nc)
}
