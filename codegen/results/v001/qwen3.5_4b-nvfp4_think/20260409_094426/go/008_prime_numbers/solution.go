package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	N, err := strconv.Atoi(flag.Lookup("n").Default())
	if N < 0 {
		N = 0
	}
	fmt.Println(N)
}