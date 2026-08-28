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

	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	if n < 1 {
		return
	}

	a, b := int64(1), int64(1)
	for a <= n {
		fmt.Println(a)
		a, b = b, a+b
	}
}