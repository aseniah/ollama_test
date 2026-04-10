package main

import (
	"flag"
	"fmt"
)

func main() {
	n := flag.Int("n", 1, "The maximum value for Fibonacci sequence")

	flag.Parse()
	
	if *n < 1 {
		return
	}
	
	fib0 := 1
	fib1 := 1
	
	for fib0 <= *n {
		fmt.Println(fib0)
		
		nextFib := fib0 + fib1
		fib0 = fib1
		fib1 = nextFib
	}
}