package main

import (
	"fmt"
)

func main(){
	num1 :=fmt.Sprintf("%b\r\n",1<<32)
	num2 :=fmt.Sprintf("%b\r\n",1<<32 - 1)
	fmt.Printf("%v\r\n",num1)
	fmt.Printf("%v\r\n",num2)

	for i:=0;i<=3 ;i++{
		for j:=0;j<5;j++{
			defer func() {
				fmt.Println(j)
			}()
			fmt.Println("\r\n")
		}
	}
}

//1 0000 0000   0000 0000  0000 0000  0000 0000