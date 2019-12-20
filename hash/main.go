package main

import (
	"fmt"
)

func main(){
	str:="i am 中国人,666!!"
	fmt.Println(len("中国人")," 中国人")
	fmt.Println("len=",len(str))
	fmt.Println("hash=",hash(str))
}

func hash(str string) uint64 {
	seed := uint64(13131)
	var hash uint64
	var temp uint64
	for i := 0; i < len(str); i++ {
		temp = uint64(str[i])
		hash = hash*seed +temp
		fmt.Printf("temp=%v	%q\r\n",temp,str[i])
	}
	fmt.Println(hash)
	return (hash & 0x7FFFFFFFFFFFFFFF)
}
