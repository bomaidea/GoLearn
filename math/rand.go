package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var random int
	random = rand.Intn(100)
	fmt.Println("My favorite number is", random)
}
