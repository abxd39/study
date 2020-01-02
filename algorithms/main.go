package main

import (
	"log"

	"github.com/abxd39/study/algorithms/sort"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	l := []int{3, 7, 9, 1, 6, 2, 5, 8, 0, 4}
	al := new(sort.Algorithms)
	al.Version1(l)
	al.Insertion(l)
	log.Println("hello algorithms !!!")
}
