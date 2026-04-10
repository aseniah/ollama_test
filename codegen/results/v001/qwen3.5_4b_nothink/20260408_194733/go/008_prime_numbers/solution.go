package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <N>")
		os.Exit(1)
	}

	n, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if n < 2 {
		return
	}

	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Println(i)
		}
	}
}

func parseInt(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return val, nil
}