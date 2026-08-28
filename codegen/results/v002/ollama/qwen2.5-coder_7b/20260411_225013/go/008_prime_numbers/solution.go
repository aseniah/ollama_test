package main

import (
	"fmt"
	"os"
)

func main() {
	n := 0
	fmt.Sscanf(os.Args[1], "%d", &n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}