package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	if n < 2 {
		return
	}

	// Sieve of Eratosthenes
	sieve := make([]bool, n+1)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	if n >= 1 {
		sieve[1] = false
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