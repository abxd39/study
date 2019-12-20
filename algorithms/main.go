package main

import (
	"log"

	"github.com/abxd39/myproject/algorithms/BubbleSort"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}
func main() {
	l := []int{3, 7, 9, 1, 6, 2, 5, 8, 0, 4}
	new(BubbleSort.BubbleSort).Version1(l)
	log.Println("hello algorithms !!!")
}
