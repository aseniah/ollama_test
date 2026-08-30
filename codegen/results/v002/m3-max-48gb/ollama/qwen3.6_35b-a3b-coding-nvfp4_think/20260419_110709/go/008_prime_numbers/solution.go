package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i <= n/i; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}