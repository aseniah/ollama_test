#!/usr/bin/env python3
import sys

def main():
    if len(sys.argv) < 2:
        print("Usage: python script.py N")
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: N must be an integer")
        return

    if n < 1:
        return

    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()