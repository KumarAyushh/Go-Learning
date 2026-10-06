package main

import "fmt"

func main() {
	//slices are the most common collection type in Go. They are more flexible than arrays and provide a powerful way to work with sequences of data. 
	// A slice is a dynamically-sized, flexible view into the elements of an array. 
	// It is a descriptor of an array segment and consists of a pointer to the array, the length of the segment, and its capacity.

	//slices are like arraylist in java
	// []type{...}
	res := []string{"ayush", "sachin", "rohit", "virat"}
	fmt.Println(res, res[0], res[len(res)-1]) // len(res)-1 gives the index of the last element

	var nums[] int

	nums = append(nums, 1)
	nums = append(nums, 2, 3)
	fmt.Println(nums)

	//slices are reference types, which means that when you assign a slice to another variable, both variables point to the same underlying array. 
	// Modifying one slice will affect the other since they share the same data.
}