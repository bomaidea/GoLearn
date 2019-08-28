package main

import "fmt"

func fibonacci() func() int {
	fibonacci := 0
	last := 0

	return func() int {
		mid := last
		last = fibonacci
		fibonacci += mid
		if fibonacci == 0 && last == 0 {
				last = 1
		}
		return fibonacci
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
