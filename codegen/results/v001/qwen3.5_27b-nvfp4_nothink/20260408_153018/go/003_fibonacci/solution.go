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

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 1 {
		return
	}

	a, b := 0, 1

	for b <= n {
		fmt.Println(b)
		a, b = b, a+b
	}
}