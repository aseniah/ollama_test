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
	if err != nil || n < 1 {
		return
	}

	a := 1
	b := 1

	for {
		if a > n {
			break
		}
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}