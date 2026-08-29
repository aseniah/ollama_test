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
	if !ok {
		return
	}

	one := big.NewInt(1)
	if n.Cmp(one) < 0 {
		return
	}

	a := big.NewInt(1)
	b := big.NewInt(1)

	for a.Cmp(n) <= 0 {
		fmt.Println(a)
		a, b = b, new(big.Int).Add(a, b)
	}
}