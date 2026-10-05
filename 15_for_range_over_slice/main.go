package main

import "fmt"

func main() {
	views := []int{100, 200, 300}
	// 🔹 Step 1: Range over the slic
	
	total := 0
	for i, v := range views {
		fmt.Println("Index:", i, "Value:", v)
		total += v
	}

	fmt.Println("Total Views:", total)
}