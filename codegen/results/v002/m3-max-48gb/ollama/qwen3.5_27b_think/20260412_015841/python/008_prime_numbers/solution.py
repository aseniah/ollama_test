import sys

def is_prime(n):
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    # Check odd divisors up to sqrt(n)
    for i in range(3, int(n ** 0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

if __name__ == "__main__":
    # Check if command line argument is provided
    if len(sys.argv) < 2:
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)
    
    # Print all primes up to and including N
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)