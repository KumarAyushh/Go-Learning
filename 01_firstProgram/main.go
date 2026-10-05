package main

//main() function is the entry point of the program
//run a go file -> go looks for package main and  the main function and executes it
//Go me code packages me organized hota hai (like folders/modules).
//Package = collection of related code files
//Jab Go program run hota hai, Go compiler sirf package main ko executable program banata hai.


//import "fmt" is used to include the fmt package, 
// which provides functions for formatted input and output such as printing text to the console.
import "fmt"

func main() {
	fmt.Println("Hello, world")
}

//Go supports utf-8 (character encoding scheme) encoding, so you can use non-ASCII characters in your code. For example, you can print "नमस्ते दुनिया" (Hello World in Hindi) using the fmt.Println function.

//we can use go build main.go to compile the code and generate an executable file. The name of the executable file will be main.exe on Windows or main on Linux/Mac. You can then run the executable file to see the output of your program.
//we can share this binary executable file with others, and they can run it without needing to have Go installed on their system because it contains all the necessary dependencies into a single file. This is one of the advantages of compiling Go code into a standalone executable.

//when we have imported a package, we can use its functions and types by prefixing them with the package name. For example, to use the Println function from the fmt package, we write fmt.Println("Hello, world"). The package name acts as a namespace to avoid naming conflicts between different packages.
//And whatever we are accessing must start with a capital letter, otherwise it will be private to the package and we won't be able to access it from outside the package. For example, if we define a function called printHello in our own package, we won't be able to access it from another package because it starts with a lowercase letter. But if we define a function called PrintHello, we can access it from other packages because it starts with an uppercase letter.
//In Go, capitalization is not just naming style—it controls package-level visibility.