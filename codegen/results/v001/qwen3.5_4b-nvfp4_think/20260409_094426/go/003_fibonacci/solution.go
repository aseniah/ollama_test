package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	var n int64
	err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || n < 1 {
		return // No output if N < 1 or invalid input
	}

	a := int64(1)
	b := int64(1)
	next := a + b

	for next <= n {
		fmt.Println(a)
		fmt.Println(next)
		a = b
		b = next
		next = a + b
	}

	if a <= n {
		fmt.Println(a)
	}
}