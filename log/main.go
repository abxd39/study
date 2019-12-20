package main

import (
	"fmt"
	"runtime"
)

func main (){
	fmt.Println("hah !!")
	_, file, line, ok := runtime.Caller(0)
	fmt.Println(ok)
	fmt.Println(file)
	fmt.Println(line)
	ff1()
}

func ff()  {
	fmt.Println("--------------------")
	_, file, line, ok := runtime.Caller(2)
	fmt.Println(ok)
	fmt.Println(file)
	fmt.Println(line)
}

func ff1()  {
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^")
	_, file, line, ok := runtime.Caller(0)
	fmt.Println(ok)
	fmt.Println(file)
	fmt.Println(line)
	ff()
}

