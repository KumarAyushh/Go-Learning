package main

import "fmt"


func main(){
	//var name type
	var name string; //this format is called variable declaration and can be used both inside and outside functions.
	name ="Go Programming Language"
	age := 21 //this format is called short variable declaration and can only be used inside functions.
	fmt.Println(name, age )

	//go uses float64 as the default type for floating-point numbers
	var rating float64 = 4.5
	fmt.Println("Rating:", rating)
}