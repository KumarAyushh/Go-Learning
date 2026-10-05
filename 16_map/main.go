package main

import "fmt"

func main() {
	//map[keyType]valueType{...}
	//maps are unordered collections of key-value pairs. They are also known as hash tables or dictionaries in other programming languages. A map is a built-in data type in Go that provides a way to associate values with keys. The keys in a map must be of a type that is comparable (e.g., strings, integers), while the values can be of any type.
	//maps are like hashmaps in java

	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}
	fmt.Println(ages, len(ages))
	//accessing values
	fmt.Println("Alice's age:", ages["Alice"])
	fmt.Println("Bob's age:", ages["Bob"])


	//make(map[K]V)
	var scores map[string]int
	fmt.Println(scores, scores["a"]);

	scores["math"] = 90
	fmt.Println(scores)
} 