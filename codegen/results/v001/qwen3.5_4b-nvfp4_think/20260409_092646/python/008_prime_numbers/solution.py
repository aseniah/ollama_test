#!/usr/bin/env python3

import sys

def is_prime(num):
    """Check if a number is prime."""
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    
    # Check only up to sqrt(num) for optimization
    for i in range(3, int(num ** 0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

def main():
    """Main function to find and print primes up to N."""
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
        # Primes only exist for n >= 2, so we can't print anything below that
        if n < 2:
            return
    except ValueError:
        print(f"Error: {sys.argv[1]} must be a valid integer")
        sys.exit(1)
    
    # Find and print all primes up to and including n
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()