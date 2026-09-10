package main

import (
	"fmt"
	"os"
	"strconv"
)

func sieve(n int) []int {
	if n < 2 {
		return nil
	}

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	primes := sieve(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}