package main

import (
	"fmt"
	"math"
	"sort"
)

func main(){

	names:=[]string{"appKey","method","signMethod","shopId","timeType","timestamp","startCreate","endCreate","v","pageNo","pageSize","messageFormat"}
	sort.Strings(names)
	for _, value := range names {
		fmt.Println(value)
	}
	var count float64
	count =-1.00000005
	fmt.Println(math.Abs(count))
}
