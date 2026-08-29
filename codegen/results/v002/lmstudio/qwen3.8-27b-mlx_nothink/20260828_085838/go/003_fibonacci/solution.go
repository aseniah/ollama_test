package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	n, ok := new(big.Int).SetString(os.Args[1], 10)
	if !ok || n.Sign() < 1 {
		return
	}

	a := big.NewInt(1)
	b := big.NewInt(1)

	for {
		if a.Cmp(n) > 0 {
			break
		}

		fmt.Println(a)

		c := new(big.Int).Add(a, b)
		a, b = b, c
	}
}