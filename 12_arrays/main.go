package main

import "fmt"

func main() {
	//size is part of the type of an array, so arrays cannot be resized.
	var marks [3]int = [3]int{10,20,30} //this right side is called array literal
	fmt.Println(marks)
	
	var updatedMarks [3]int

	updatedMarks[0] = 40
	updatedMarks[1] = 50
	updatedMarks[2] = 60	
	fmt.Println(updatedMarks)


	//array literal
	//“An array literal is a way to declare and initialize an array in a single expression using values enclosed in curly braces.”
	res := [5]int{2,3,4,5,6} //this right side is called array literal
	fmt.Println(res)
	fmt.Println(len(res))
	

}