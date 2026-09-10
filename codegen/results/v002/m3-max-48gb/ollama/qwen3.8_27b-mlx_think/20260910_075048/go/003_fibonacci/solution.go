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

	N, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	if N < 1 {
		return
	}

	a, b := int64(1), int64(1)
	for a <= N {
		fmt.Println(a)
		a, b = b, a+b
	}
}