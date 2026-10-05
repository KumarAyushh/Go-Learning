package main;

import (
	"fmt"
	"strings"
)

func main() {
	firstName := "John"
	lastName := "Doe"

	fullName := firstName + lastName //string concatenation
	fmt.Println("Full Name:", fullName)
	fmt.Println(strings.ToUpper(fullName))
}