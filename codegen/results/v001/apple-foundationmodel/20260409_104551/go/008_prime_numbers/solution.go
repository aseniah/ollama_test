package main

import (
    "fmt"
    "os"
)

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))) + 1; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <N>")
        return
    }

    N, _ := os.Atoi(os.Args[1])

    for i := 2; i <= N; i++ {
        if isPrime(i) {
            fmt.Println(i)
        }
    }
}