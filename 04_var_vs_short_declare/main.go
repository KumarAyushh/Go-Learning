package main

import "fmt"

func main() {
	var city string
	city = "New York"
	//fmt.Println(city)

	var name = "John Doe" //infer type from the value, inferred to string

	// := is a shorthand operator in Go used to declare and initialize variables inside functions without explicitly specifying the data type. 
	// Go automatically infers the variable type from the assigned value.

	subscribers := 1000 //declare and initialize in one line, type inferred to int

	likes, comments := 50, 100; //multiple variables declare and initialize in one line
	fmt.Println("City:", city)
	fmt.Println("Name:", name)
	fmt.Println("Subscribers:", subscribers)
	fmt.Println("Likes:", likes, "Comments:", comments)
}