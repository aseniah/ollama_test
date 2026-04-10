#!/usr/bin/env python3
import sys

if __name__ == "__main__":
    # Validate that exactly one argument was provided
    if len(sys.argv) != 2:
        sys.exit(1)

    # Get N as an integer
    try:
        N = int(sys.argv[1])
    except ValueError:
        sys.exit(1)

    # Exit if N is less than 1
    if N < 1:
        sys.exit(0)

    # Generate and print Fibonacci numbers up to N
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b