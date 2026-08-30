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

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	if N < 1 {
		return
	}

	a := 1
	b := 1

	fmt.Println(a)
	if N >= b {
		fmt.Println(b)
	}

	for {
		c := a + b
		if c > N {
			break
		}
		fmt.Println(c)
		a = b
		b = c
	}
}