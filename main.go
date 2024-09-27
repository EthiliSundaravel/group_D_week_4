// main.go
package main

import (
	"fmt"
)

func main() {

	
	fmt.Println("Welcome to Group A's Week 4 Project!")
	
    var num int
    var operation string

    fmt.Print("Enter the number: ")
    fmt.Scanln(&num)
    

    fmt.Println("Choose an operation : \n 1 : count digits in a number \n 2 : check for palindrome \n 3 : check for prime number \n 4 : factorial of a number")
    fmt.Scanln(&operation)

    switch operation {
    case "1":
        fmt.Printf("total digits in a number %d = %d\n", num, countDigits(num))
    
    default:
        fmt.Println("Invalid operation! Please choose from 1 to 4")
    }
}
