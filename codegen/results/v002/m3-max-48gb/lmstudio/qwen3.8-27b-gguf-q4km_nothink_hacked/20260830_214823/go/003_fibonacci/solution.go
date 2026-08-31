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
	
	nStr := os.Args[1]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return
	}
	
	if n < 1 {
		return
	}
	
	a, b := 1, 1
	// Print first 1
	if a <= n {
		fmt.Println(a)
	}
	if b <= n && a == b {
		fmt.Println(b)
	}
	
	a, b = 1, 2
	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}