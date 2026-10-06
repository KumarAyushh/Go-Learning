package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	//implicit conversion
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
}