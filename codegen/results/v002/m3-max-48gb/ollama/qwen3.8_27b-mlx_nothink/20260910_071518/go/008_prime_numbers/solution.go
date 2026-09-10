package main

import (
	"fmt"
	"os"
	"strconv"
)

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

	var primes []int
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	primes := sieveOfEratosthenes(n)
	for _, p := range primes {
		fmt.Println(p)
	}
}