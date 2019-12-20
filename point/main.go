package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "W  Yw  "
	str = compressStr(str)
	defer fmt.Println(str)
}

//利用正则表达式压缩字符串，去除空格或制表符
func compressStr(str string) string {
	if str == "" {
		return ""
	}
	// 匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}
