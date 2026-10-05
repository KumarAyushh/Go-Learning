package main

import "fmt"

func main() {

    // variable declaration (short syntax)
    day := 5

    // switch checks the value of 'day'
    switch day {

    // if day == 1
    case 1:
        fmt.Println("Monday")

    // if day == 2
    case 2:
        fmt.Println("Tuesday")

    // if day == 3
    case 3:
        fmt.Println("Wednesday")

    // if day == 4
    case 4:
        fmt.Println("Thursday")

    // if day == 5
    case 5:
        fmt.Println("Friday")

    // runs when none of the above cases match
    default:
        fmt.Println("Weekend")
    }

}