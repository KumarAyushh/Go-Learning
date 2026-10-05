package main

import "fmt"

// func main() {
// 	fmt.Println("Addition is : ", add(5, 10)) //we use comma to separate the arguments passed to the Println function.
// }

// func add(x int, y int) int {
// 	return x + y
// }



//In Go, we can define a function that takes multiple parameters of the same type by 
// specifying the type only once after the last parameter. For example, instead of writing `func add(x int, y int) int`, 
// we can write `func add(x, y int) int`. This makes the code more concise and easier to read.
//In Go, we can define a function that returns multiple values by specifying the types of 
// the return values in parentheses after the function signature. For example, 
// we can define a function `swap` that takes two string parameters and returns two string values as follows:


// func swap(x, y string) (string, string) {
	
// 	return y, x
// }

// func main() {
// 	a, b := swap("Hello", "World")
// 	fmt.Println(a, b)
// }


func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return //we can omit the return values in the return statement when we have named return values. In this case, the function will return the current values of the named return variables x and y.
}

func main() {
	fmt.Println(split(17))
}

//when we are using named return values
//then those variables are initialized to their zero values (0 for int, "" for string, false for bool, etc.) when the function is called.
//so in the above code before the start of the funciton
//two functions are implicitly created x and y of type int and initialized to 0