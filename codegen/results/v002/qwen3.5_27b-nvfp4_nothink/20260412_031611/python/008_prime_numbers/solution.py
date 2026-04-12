import sys

def sieve_primes(n: int):
    """Generate all prime numbers up to and including n using the Sieve of Eratosthenes."""
    if n < 2:
        return
    
    # Boolean list where index represents number; True means potentially prime
    is_prime = [True] * (n + 1)
    is_prime[0] = False
    is_prime[1] = False

    p = 2
    while p * p <= n:
        if is_prime[p]:
            # Mark multiples of p starting from p*p
            for i in range(p * p, n + 1, p):
                is_prime[i] = False
        p += 1
    
    # Collect and print primes
    for num in range(2, n + 1):
        if is_prime[num]:
            print(num)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)
    
    sieve_primes(n)