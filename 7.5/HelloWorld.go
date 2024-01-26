package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		isPrime := true
		if i == 2 {
			fmt.Println(i)
		} else if i%2 == 0 {
			continue
		} else {
			for a := 3; a*a <= i; a++ {
				if i%a == 0 {
					isPrime = false
					break
				}
			}
			if isPrime == true {
				fmt.Println(i)
			}
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
