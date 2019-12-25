package main

import (
	"log"

	"github.com/abxd39/study/algorithms/sort"

	_ "github.com/abxd39/study/algorithms/sort"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}
func main() {
	l := []int{3, 7, 9, 1, 6, 2, 5, 8, 0, 4}
	//new(sort.Algorithms).Version1(l)
	new(sort.Algorithms).Insertion(l)
	log.Println("hello algorithms !!!")
}
