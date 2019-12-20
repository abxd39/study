package main

import (
	"fmt"
	"unicode"
)

func main(){
	fmt.Println("\u7981")
	fmt.Println("\u6709")
	var run rune
	run= 1111

	b:=unicode.IsPrint(run)
	fmt.Println(b)
}
