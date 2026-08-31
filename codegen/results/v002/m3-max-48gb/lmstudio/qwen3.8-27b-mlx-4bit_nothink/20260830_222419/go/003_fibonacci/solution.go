package main

import (
	"fmt"
	"strconv"
	"os"
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

	a := 1
	b := 1
	if n >= 1 {
		fmt.Println(a)
		if n >= 1 && b <= n {
			fmt.Println(b)
		}
	}
	c := a + b
	for c <= n {
		fmt.Println(c)
		a = b
		b = c
		c = a + b
	}
}