package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1} // Y default 0
	v3 = Vertex{} // X, Y default 0
	p = &Vertex{1,2} // has type *
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
