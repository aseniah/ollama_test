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

	if n >= 1 {
		fmt.Println(1)
	}
	if n >= 1 {
		fmt.Println(1)
	}

	a, b := 1, 1
	for {
		next := a + b
		if next > n {
			break
		}
		fmt.Println(next)
		a, b = b, next
	}
}