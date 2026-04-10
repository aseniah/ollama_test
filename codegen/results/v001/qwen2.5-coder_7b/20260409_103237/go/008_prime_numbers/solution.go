package main

import (
	"fmt"
	"os"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run primes.go <N>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		os.Exit(1)
	}

	primesFound := false
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
			primesFound = true
		}
	}

	if !primesFound {
		fmt.Println("No primes found up to", n)
	}
}