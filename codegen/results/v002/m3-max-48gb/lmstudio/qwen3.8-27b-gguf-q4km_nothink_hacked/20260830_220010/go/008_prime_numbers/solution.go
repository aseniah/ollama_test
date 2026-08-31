package main

import (
	"fmt"
	"os"
	"strconv"
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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: program N")
		os.Exit(1)
	}
	
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: N must be an integer")
		os.Exit(1)
	}
	
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}