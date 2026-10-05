package main;

import (
	"fmt"
)

func main() {
	views1 := 1000
	views2 := 2000

	totalViews := views1 + views2

	likes := 10;
	likes ++;

	avgViews := totalViews/2

	fmt.Println("Total Views:", totalViews)
	fmt.Println("Average Views:", avgViews)
	fmt.Println("Likes:", likes)
}