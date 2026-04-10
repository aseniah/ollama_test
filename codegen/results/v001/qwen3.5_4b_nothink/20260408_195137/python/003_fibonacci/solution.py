#!/usr/bin/env python3
import sys

def main():
    # Read N from command line argument
    if len(sys.argv) != 2:
        print(f"Usage: {sys.argv[0]} N")
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        print(f"Invalid input: {sys.argv[1]} is not a valid integer")
        sys.exit(1)

    # If N < 1, print nothing
    if n < 1:
        return

    # Start generating Fibonacci numbers
    fib1 = 1
    fib2 = 1

    # We have already printed the first 1 (since sequence starts 1, 1, 2...)
    # But to be clear based on "Start the sequence: 1, 1, 2, 3, 5", 
    # we print 1, then calculate next, print if <= N.
    
    # Print the first 1
    print(fib1)
    
    while True:
        # Calculate next Fibonacci number
        fib3 = fib1 + fib2
        fib1 = fib2
        fib2 = fib3
        
        # Print if it does not exceed N
        if fib3 <= n:
            print(fib3)
        else:
            break

if __name__ == "__main__":
    main()