package main

import "math"


func isPrimeNumber(num int) bool{
	if num <= 1 {
		return false // 0 and 1 are not prime
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false // Not prime
		}
	}
	return true // Is prime
}