package main

import (
	"log"
	"strings"

	"github.com/abxd39/study/algorithms/sort"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	l := []int{3, 7, 9, 1, 6, 2, 5, 8, 0, 4}
	al := new(sort.Algorithms)
	al.Version1(l)
	ll := strings.Split(" 100,105,106,111,116,120,122,126,17,19,20,21,22,23,26,27,28,29,30,31,41,46,48,49,50,51,53,54,55,59,62,63,64,67,68,69,72,73,75,76,77,78,87,89,94,98", ",")
	al.Insertion(ll)
	log.Println(strings.Compare("126", "121"))
	log.Println("hello algorithms !!!")
}
