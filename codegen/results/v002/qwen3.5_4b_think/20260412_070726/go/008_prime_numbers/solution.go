package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n % 2 == 0 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: program <N>")
		os.Exit(1)
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid input: %s\n", err)
		os.Exit(1)
	}

	if N < 2 {
		return
	}

	for i := 2; i <= N; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}