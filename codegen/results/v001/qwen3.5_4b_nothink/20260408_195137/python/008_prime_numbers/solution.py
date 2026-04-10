import sys

def has_prime_up_to(n):
    """Check if there is at least one prime up to n (inclusive)."""
    if n < 2:
        return False
    # 2 is prime
    return True

def print_primes_up_to(n):
    """Print all prime numbers up to and including n, one per line."""
    if n < 2:
        return
    if n >= 2:
        print(2)
    # Check odd numbers from 3 to n
    i = 3
    while i <= n:
        is_prime = True
        # Only check up to sqrt(i)
        limit = int(i ** 0.5)
        if limit >= 2:
            for j in range(3, limit + 1, 2):
                if i % j == 0:
                    is_prime = False
                    break
        if is_prime:
            print(i)
        i += 2

if len(sys.argv) != 2:
    sys.exit(1)

try:
    n = int(sys.argv[1])
    if n < 1:
        # No primes for n < 2
        sys.exit(0)
    print_primes_up_to(n)
except ValueError:
    sys.exit(1)