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
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprint(os.Stdout, "Usage: program <N>\n")
		os.Exit(1)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil || n < 0 {
		fmt.Fprintln(os.Stderr, "Error: invalid integer argument N")
		os.Exit(1)
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}