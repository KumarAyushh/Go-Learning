package main

import "fmt"

func main() {
	i,j := 42, 2701

	p := &i         // point to the address of i
	fmt.Println(p)  // print the address of i
	fmt.Println(*p) // read i through the pointer

	*p = 21		 // set i through the pointer
	fmt.Println(*p) // read the new value of i

	p = &j         // point to the address of j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(*p) // read the new value of j
}