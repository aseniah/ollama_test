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
	if err != nil {
		return
	}

	if n < 2 {
		return
	}

	// Use Sieve of Eratosthenes to find all primes up to N
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if sieve[i] {
			fmt.Println(i)
		}
	}
}