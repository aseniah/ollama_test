import sys
import math

def sieve(n):
    """Return a list of all primes up to and including n using Sieve of Eratosthenes."""
    if n < 2:
        return []
    
    is_prime = [True] * (n + 1)
    is_prime[0] = False
    is_prime[1] = False
    
    for i in range(2, int(math.isqrt(n)) + 1):
        if is_prime[i]:
            for j in range(i * i, n + 1, i):
                is_prime[j] = False
    
    return [i for i in range(2, n + 1) if is_prime[i]]

def main():
    # Read the single integer argument from command line
    try:
        if len(sys.argv) < 2:
            # No argument provided, print nothing
            return
        n = int(sys.argv[1])
    except ValueError:
        return
    
    primes = sieve(n)
    for p in primes:
        print(p)

if __name__ == "__main__":
    main()