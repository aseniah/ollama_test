package main

import (
	"flag"
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	N := flag.Int("N", -1, "The upper limit for prime numbers")
	flag.Parse()

	*flag.SetOutput = os.Stdout // Just to ensure output goes to stdout if needed (safety)
	fmt.Println(*N) // This is wrong. Let's fix this in the actual logic.
}