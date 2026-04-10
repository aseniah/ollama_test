import sys
import math

def sieve(n):
    if n < 2:
        return
    
    # Create a boolean array "is_prime[0..n]" and initialize
    # all entries it as true. A value in is_prime[i] will
    # finally be false if i is Not a prime, else true.
    is_prime = [True] * (n + 1)
    is_prime[0] = False
    is_prime[1] = False
    
    for p in range(2, int(math.sqrt(n)) + 1):
        if is_prime[p]:
            # Update all multiples of p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
    
    # Print primes found
    for p in range(2, n + 1):
        if is_prime[p]:
            print(p)

def main():
    if len(sys.argv) != 2:
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    sieve(n)

if __name__ == "__main__":
    main()