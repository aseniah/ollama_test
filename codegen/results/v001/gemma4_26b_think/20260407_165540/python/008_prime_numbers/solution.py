#!/usr/bin/env python3
import sys

def print_primes_up_to(n):
    """
    Uses the Sieve of Eratosthenes algorithm to find and print 
    all prime numbers up to and including n.
    """
    if n < 2:
        return

    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as true. A value in is_prime[i] will
    # finally be false if i is Not a prime, else true.
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False

    for p in range(2, int(n**0.5) + 1):
        # If is_prime[p] is not changed, then it is a prime
        if is_prime[p]:
            # Updating all multiples of p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False

    # Print all prime numbers
    for p in range(2, n + 1):
        if is_prime[p]:
            print(p)

def main():
    # Check if exactly one argument is provided (excluding the script name)
    if len(sys.argv) != 2:
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
        print_primes_up_to(n)
    except ValueError:
        # If the argument is not a valid integer, do nothing
        return

if __name__ == "__main__":
    main()