package main

import (
	"fmt"
	"github.com/abxd39/myproject/pag/packageC"
)

func main(){
	fmt.Println("Hello 世界！！")
	new(packageC.CTemp).GetName()
	fmt.Println(f(3))
}


func f(n int)(r int){
	defer func() {
		fmt.Println(r)
		r+=n
		err:=recover()
		if err!=nil{
			fmt.Println(err)
		}
	}()
	var f func()
	defer  f()
	f= func() {
		r+=2
	}
	fmt.Println(r)
	return n+1
}
