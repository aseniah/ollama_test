import sys

def sieve_of_eratosthenes(n):
    """Generate all prime numbers up to and including n using Sieve of Eratosthenes"""
    if n < 2:
        return []
    
    # Initialize boolean array
    is_prime = [True] * (n + 1)
    is_prime[0] = is_prime[1] = False
    
    for i in range(2, int(n**0.5) + 1):
        if is_prime[i]:
            for j in range(i*i, n + 1, i):
                is_prime[j] = False
    
    return [i for i in range(2, n + 1) if is_prime[i]]

def main():
    if len(sys.argv) != 2:
        print("Usage: python program.py N", file=sys.stderr)
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: N must be an integer", file=sys.stderr)
        sys.exit(1)
    
    primes = sieve_of_eratosthenes(n)
    
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()