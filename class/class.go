package class

import ( 
	"fmt"
	"time"
)

func Palindrome(a string) bool {
	n := len(a)

    for i := 0; i < n/2; i++ {
        if a[i] != a[n-1-i] {
            return false
        }
    }
    return true
}
func Fibinocci(x int) int {
	c, b := 0, 1
	for i := 0; i < x; i++ {
		fmt.Println(c)
		c, b = b, c+b
	}
     return x
}
func Factorial(n int)int{
	if n ==0{
		return 1
	}
	return n*Factorial(n-1)
}
func Greeting(lang, msg string) {
		fmt.Printf("[%s]: %s\n", lang, msg)
		time.Sleep(1 * time.Second)
}