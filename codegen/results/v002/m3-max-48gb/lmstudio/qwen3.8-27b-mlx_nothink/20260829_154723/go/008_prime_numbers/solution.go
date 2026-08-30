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

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

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