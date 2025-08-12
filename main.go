package main

import (
    "fmt"
    "go-demo/utils" // Adjust module name or use relative import if needed
)

func main() {
    // Basic data types
    var integer int = 42
    var floating float64 = 3.14
    var boolean bool = true
    var str string = "Hello, Go!"

    fmt.Println("Integer:", integer)
    fmt.Println("Float:", floating)
    fmt.Println("Boolean:", boolean)
    fmt.Println("String:", str)

    // Using flow control from utils package
    fmt.Println("\nChecking number flow control with 15:")
    utils.CheckNumber(15)

    fmt.Println("\nChecking number flow control with -5:")
    utils.CheckNumber(-5)

    fmt.Println("\nLoop example with n=5:")
    utils.LoopExample(5)

    // Demonstrate simple for loop in main
    fmt.Println("\nSimple for loop in main:")
    for i := 1; i <= 3; i++ {
        fmt.Printf("Iteration %d\n", i)
    }

    // Defer example directly in main
    defer fmt.Println("Deferred statement: This prints at the end of main")

    fmt.Println("Main function is about to end")
}
