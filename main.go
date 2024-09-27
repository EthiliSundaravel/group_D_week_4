// main.go
package main

import (
	"fmt"
)

func main() {

	
	fmt.Println("Welcome to Group A's Week 4 Project!")
	
    var num int
    var operation string

    fmt.Println("Choose an operation : \n 1 : count digits in a number \n 2 : check for palindrome \n 3 : check for prime number \n 4 : factorial of a number")
    fmt.Scanln(&operation)

    fmt.Print("Enter the number: ")
    fmt.Scanln(&num)

    switch operation {
    case "1":
        fmt.Printf("total digits in a number %d = %d\n", num, countDigits(num))
    
    case "3":
        if isPrimeNumber(num) {
            fmt.Printf("%d is a prime number.\n", num)
        } else {
            fmt.Printf("%d is not a prime number.\n", num)
        }
    case "4":
        fmt.Printf("factorial of a number %d is %d\n", num, factorial(num))

    
    default:
        fmt.Println("Invalid operation! Please choose from 1 to 4")
    }
}
