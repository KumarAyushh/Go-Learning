package main

import "fmt"


func main() {

	//untyped constant
	const appName = "Go Basics"

	//typed constant
	const maxUpload int = 25


	const discountedPrice  float64 = 10.3

	fmt.Println("App Name:", appName)
	fmt.Println("Max Upload Size:", maxUpload, "MB")
	fmt.Println("Discounted Price:", discountedPrice, "USD")	
}
