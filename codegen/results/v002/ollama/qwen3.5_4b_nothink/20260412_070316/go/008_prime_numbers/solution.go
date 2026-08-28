package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 2, "The upper bound for prime numbers")
	flag.Parse()

	if *n < 2 {
		return
	}

	isPrime := make([]bool, *n+1)
	isPrime[0] = false
	isPrime[1] = false

	// Sieve of Eratosthenes
	for i := 2; i*i <= *n; i++ {
		if isPrime[i] {
			for j := i * i; j <= *n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Print primes
	for i := 2; i <= *n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}