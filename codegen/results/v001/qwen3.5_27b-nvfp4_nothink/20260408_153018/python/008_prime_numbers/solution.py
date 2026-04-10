import sys
import math

def sieve_of_eratosthenes(limit):
    if limit < 2:
        return []
    
    # Create a boolean array "is_prime[0..limit]" and initialize all entries as true.
    is_prime = [True] * (limit + 1)
    is_prime[0] = is_prime[1] = False
    
    for p in range(2, int(math.isqrt(limit)) + 1):
        if is_prime[p]:
            # Mark multiples of p starting from p*p as not prime
            for i in range(p * p, limit + 1, p):
                is_prime[i] = False
    
    return [i for i in range(2, limit + 1) if is_prime[i]]

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Argument must be an integer.")
        sys.exit(1)

    primes = sieve_of_eratosthenes(n)
    
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()