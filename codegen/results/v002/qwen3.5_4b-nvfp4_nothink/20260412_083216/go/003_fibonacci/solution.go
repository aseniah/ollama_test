package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	
	// Read the integer argument N from stdin as per Go convention if arg[0] is empty or not provided.
	// The prompt says "Accept a single integer argument N from the command line", 
	// but in Go we typically read from os.Args[1]. Let's check for existence first.
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Please provide an integer argument")
		os.Exit(1)
	}

	n, err := fmt.Scanf("%d", nil) // This won't work directly in standard input reading logic for args.
	// Let's implement proper argument parsing instead of scanf (which reads from stdin).
	// Re-doing: Use fmt.Scanf which requires an actual stream or use strconv.ParseInt
	
	fmt.Sscanf(os.Args[1], "%d", &n) // Correct approach? No, let's do it properly.

}