package main

import "fmt"

func main()  {
	multiParam("wang","ying","wen")
}


func multiParam(args...string){
	for index,v:=range args{
		fmt.Printf("下标为：--->%d值为--->%v\r\n\n",index,v)
	}
}

