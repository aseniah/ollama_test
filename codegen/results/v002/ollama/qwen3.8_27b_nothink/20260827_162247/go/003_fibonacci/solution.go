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
	
	if n < 1 {
		return
	}
	
	// Print Fibonacci numbers up to and including the largest one that doesn't exceed N
	a, b := 1, 1
	for {
		if a > n {
			break
		}
		fmt.Println(a)
		a, b = b, a+b
	}
}