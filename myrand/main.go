package main

import (
	"fmt"

	"github.com/lifei6671/gorand"
)

func main() {
	fmt.Println(gorand.RandomAlphabetic(16))
	fmt.Println(string(gorand.KRand(20, gorand.KC_RAND_KIND_ALL)))
}
