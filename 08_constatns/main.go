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

// 🔥 Quick Comparison
// Feature	Untyped Constant	Typed Constant
// Type specified?	❌ No	✅ Yes
// Flexibility	High	Low
// Type decided	Usage ke time	Declaration ke time
// Example	const x = 10	const x int = 10
// 🧠 Easy Memory Trick
// Untyped = Universal
// Typed = Tight restriction
// Real-world usage