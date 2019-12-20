package main

import (
	"log"
)

func init()  {
	log.SetFlags(log.Lshortfile|log.Ldate|log.Ltime)
}

//go test -bench=. -benchmem -run=none
//go tool compile -S *.go
func main(){
	log.Println(fn(3))
}


func fn(n int)(r int){
	defer func() {
		log.Println(".....")
		r+=n
		log.Println(recover())
	}()
	var f func()
	defer f()
	f= func() {
		r+=2
	}
	f1:=func()int {
		log.Println("return")
		return n+1
	}
	return f1()
}