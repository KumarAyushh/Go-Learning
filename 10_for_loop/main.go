package main

import (
	"fmt"
)	

func main() {
	for i:= 1; i <=5; i++ {
		fmt.Println("Iteration:", i)
	}

	M:= 10

	sum := 0

	for i:=1; i <=M; i++ {
		sum += i;
	}
	fmt.Println("Sum:", sum)
}