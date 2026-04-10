package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		return
	}

	sieve(n)
}

// sieve uses the Sieve of Eratosthenes to find and print all primes up to n.
func sieve(limit int) {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}