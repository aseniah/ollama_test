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
	if err != nil {
		return
	}

	if n < 2 {
		return
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}