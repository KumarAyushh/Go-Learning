package main

import "fmt"

type Vert struct{
	X, Y int
}

var (
	v1 = Vert{1, 2}  // has type Vert
	v2 = Vert{X: 1}  // Y:0 is implicit
	v3 = Vert{}      // X:0 and Y:0 are implicit
	p  = &Vert{1, 2} // has type *Vert
)

func StructDemo() {
	fmt.Println(v1, p, v2, v3)
}