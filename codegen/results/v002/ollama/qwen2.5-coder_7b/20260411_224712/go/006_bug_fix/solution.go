package main

import "fmt"

func fizzbuzz(n int) []string {
    result := []string{}
    for i := 1; i <= n { // Fixed: start from 1 and iterate up to n inclusive
        if i%15 == 0 {
            result = append(result, "FizzBuzz")
        } else if i%3 == 0 {
            result = append(result, "Fizz")
        } else if i%5 == 0 {
            result = append(result, "Buzz")
        } else {
            result = append(result, fmt.Sprintf("%d", i))
        }
    }
    return result
}

func main() {
    for _, v := range fizzbuzz(15) {
        fmt.Println(v)
    }
}