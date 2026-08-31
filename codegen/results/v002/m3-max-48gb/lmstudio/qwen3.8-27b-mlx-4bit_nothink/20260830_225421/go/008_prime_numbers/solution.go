package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <N>\n", os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: invalid integer: %s\n", os.Args[1])
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
	sieve[1] = false

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