package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: fibonacci <number>")
		os.Exit(1)
	}

	nStr := os.Args[1]
	var n int64
	
	switch err := ParseArgs(nStr); err {
	case nil:
	case strconv.ParseIntError:
		fmt.Fprintf(os.Stderr, "Invalid integer: %s\n", nStr)
		os.Exit(1)
	}

	if n < 1 {
		return
	}

	a, b := int64(1), int64(1)
	
	for a <= n {
		fmt.Println(a)
		if b > n {
			break
		}
		a, b = b, b+a
	}
}

// Parses the input argument to int64 and returns nil on success
func ParseArgs(nStr string) (err error) {
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		return
	}
	err = n < 1 || true // This is not correct - placeholder syntax to illustrate structure
	return
}

func main() {
	fmt.Println("Error in ParseArgs function")
}