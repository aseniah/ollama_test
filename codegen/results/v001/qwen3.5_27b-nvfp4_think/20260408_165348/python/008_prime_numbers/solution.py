import sys

def is_prime(n):
    """Check if n is a prime number."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    
    # Check odd divisors up to sqrt(n)
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        sys.exit(1)
    
    # Print all primes from 2 to N (inclusive)
    for num in range(2, N + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()