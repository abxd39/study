package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var ml []int
var r *rand.Rand

func main() {
	//arrayStudy()
	//sliceStudy()
	//newAndMakeIsDifferent()
	//sliceCopy()
	Al200Algorithms()
}

func newAndMakeIsDifferent() {
	fmt.Println("*******切new 与make 的区别**********")
	ml = make([]int, 0, 10)
	fmt.Println("1_len=", len(ml))
	fmt.Println("1_cap=", cap(ml))
	//s1:=make([]int,0,10)
	type Tes struct {
		s1 []int
	}
	tes := new(Tes)
	//tes.
	fmt.Println("-----s1_len=", len(tes.s1))
	fmt.Println("-----s1_cap=", cap(tes.s1))
	tes.s1 = append(tes.s1, 0)
	tes.s1 = append(tes.s1, 1)
	tes.s1 = append(tes.s1, 2)
	tes.s1 = append(tes.s1, 3)
	fmt.Println("*****s1_len=", len(tes.s1))
	fmt.Println("*****s1_cap=", cap(tes.s1))
	err := os.Setenv("wangyingwen", "E:\\WorkSpace\\src\\document")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Format("20060102150405"))
}

func sliceStudy() {
	slice := make([]int, 0, 100)
	shopIdList := []int{726, 654, 720, 639, 725, 706, 705, 629, 731, 636, 658, 673, 749, 679, 637, 739, 633, 683, 723, 670, 747, 665, 733, 652}
	for index, value := range shopIdList {
		fmt.Printf("index=%v,value=%v\r\n", index, value)
	}

	fmt.Printf("cap=%v\r\n", cap(slice))
	fmt.Printf("cap=%v\r\n", len(slice))
	fmt.Println("slice....................")
	slice = append(slice[:len(slice):len(slice)], 1)
	slice = append(slice, 2)
	fmt.Println(len(slice))
	fmt.Println(slice)
	m := map[string]int{"a": 97, "b": 98}
	for key, value := range m {
		fmt.Printf("key=%v,value=%v\r\n", key, value)
	}
}

func arrayStudy() {
	for i := range (*[5]int)(nil) {
		//for i,v:=range (*[5]int)(nil){ //不能取值。
		v := 1
		log.Printf("index=%v,value=%v", i, v)
	}
}

/*
切片 赋值与使用copy的区别
使用 copy
 如果目标 切片大于 源切片，如果目标切片长，则不会被清空。
使用赋值
会清空 详情请执行以下代码
*/

func SliceCopy() {
	shopIdList := []int{726, 654, 720, 639, 725, 706, 705, 629, 731, 636, 658, 673, 749, 679, 637, 739, 633, 683, 723, 670, 747, 665, 733, 652}
	//var s1= shopIdList[1:6]
	var s2 = shopIdList[1:9]
	//log.Println("s1[1:6]",s1)
	log.Println("s2[1:9]", s2)
	copy(s2, shopIdList[9:13])
	log.Println("copy s2[9:13]", s2)
	s2 = shopIdList[9:13]
	//log.Println("copy s1[9:13]",s1)
	log.Println("= s2[9:13]", s2)
	copy(s2, shopIdList[8:9])
	log.Println("copy s2[8:9]", s2)

}

/*
切片以200 个为最大数量 切分的算法实现
分别用 大于200 等于200 小于200 的数测试
*/

func Al200Algorithms() {
	tlen := 201
	count := tlen / 200
	cc := tlen % 200
	if cc != 0 {
		count += 1
	}
	for i := 0; i < count; i++ {
		last := 0
		bengin := 0
		if tlen > 200 {
			last = i*200 + 200
			if last > tlen {
				last = tlen
			}
			bengin = i * 200
			log.Printf("begin=%d,last=%d", bengin, last)
		} else {
			log.Printf("begin=%d,last=%d", 0, tlen)
		}

	}
}
