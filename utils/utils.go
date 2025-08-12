package utils

import "fmt"

// CheckNumber demonstrates if-else and switch with flow control
func CheckNumber(num int) {
    if num < 0 {
        fmt.Println("Number is negative")
    } else if num == 0 {
        fmt.Println("Number is zero")
    } else {
        fmt.Println("Number is positive")
    }

    switch {
    case num < 10:
        fmt.Println("Number is less than 10")
    case num >= 10 && num < 100:
        fmt.Println("Number is between 10 and 99")
    default:
        fmt.Println("Number is 100 or more")
    }
}

// LoopExample demonstrates for loops and defer
func LoopExample(n int) {
    fmt.Println("For loop from 0 to", n-1)
    for i := 0; i < n; i++ {
        fmt.Println(i)
    }

    fmt.Println("Using defer to print numbers in reverse:")
    for i := 0; i < n; i++ {
        defer fmt.Println(i)
    }
}
