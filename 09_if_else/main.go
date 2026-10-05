package main

import "fmt"

func main() {

    // 1️⃣ Basic if
    age := 18

    if age >= 18 {
        fmt.Println("Adult")
    }


    // 2️⃣ if-else
    marks := 40

    if marks >= 50 {
        fmt.Println("Pass")
    } else {
        fmt.Println("Fail")
    }


    // 3️⃣ if - else if - else chain
    score := 82

    if score >= 90 {
        fmt.Println("Grade A")
    } else if score >= 75 {
        fmt.Println("Grade B")
    } else if score >= 50 {
        fmt.Println("Grade C")
    } else {
        fmt.Println("Fail")
    }


    // 4️⃣ if with short variable declaration (very common in Go)
    if num := 10; num%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // note:
    // variable 'num' only available inside this if-else block


    // 5️⃣ Multiple conditions (AND, OR)
    temperature := 30
    isSunny := true

    if temperature > 25 && isSunny {
        fmt.Println("Go outside")
    }

    if temperature < 20 || !isSunny {
        fmt.Println("Stay inside")
    }


    // 6️⃣ Nested if
    username := "admin"
    password := "1234"

    if username == "admin" {
        if password == "1234" {
            fmt.Println("Login success")
        } else {
            fmt.Println("Wrong password")
        }
    }


    // 7️⃣ No parentheses needed (IMPORTANT difference from Java/JS)
    num2 := 5

    if num2 > 0 {
        fmt.Println("Positive")
    }


    // 8️⃣ Condition must be boolean (no automatic conversion)
    // ❌ invalid in Go
    // if 1 { }

    flag := true
    if flag {
        fmt.Println("Condition true")
    }


    // 9️⃣ Early return style (common in Go functions)
    number := -1

    if number < 0 {
        fmt.Println("Negative number")
        return
    }

    fmt.Println("Positive number")


    // 🔟 Comparing strings
    name := "Ayush"

    if name == "Ayush" {
        fmt.Println("Hello Ayush")
    }

}