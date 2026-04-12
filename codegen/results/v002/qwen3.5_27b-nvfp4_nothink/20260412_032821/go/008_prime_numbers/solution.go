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

	primes := sieveOfEratosthenes(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}

// sieveOfEratosthenes returns a slice of all prime numbers up to and including limit.
func sieveOfEratosthenes(limit int) []int {
	if limit < 2 {
		return nil
	}

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

	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}