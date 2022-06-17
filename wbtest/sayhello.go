package main

import (
	"fmt"
)

//testgit3
func isPrime(number int) bool {
	isPrime := true
	if number >= 0 {
		if number == 0 || number == 1 {
			isPrime = false
			fmt.Println(isPrime)
		} else {
			for i := 2; i <= number/2; i++ {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime == true {
				fmt.Println(isPrime)
			}
		}
	} else {
		isPrime = false
		fmt.Println(isPrime)
	}
	return isPrime
}
func main() {
	var number int
	fmt.Print("Enter Number To check Prime:")
	fmt.Scanln(&number)
	isPrime(number)
}
