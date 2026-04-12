package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Get the command line argument
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: go run fibonacci.go <N>")
        os.Exit(1)
    }

    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error: N must be a valid integer")
        os.Exit(1)
    }

    // If N < 1, print nothing
    if n < 1 {
        return
    }

    // Print the Fibonacci sequence up to N
    a, b := 1, 1
    for a <= n {
        fmt.Println(a)
        a, b = b, a+b
    }
}