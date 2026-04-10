import sys

def is_prime(n):
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    # Get N from command line arguments
    if len(sys.argv) != 2:
        print("Usage: python script.py N", file=sys.stderr)
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print(f"Error: Expected an integer, got '{sys.argv[1]}'", file=sys.stderr)
        sys.exit(1)
    
    # Ensure N is non-negative
    if n < 0:
        return

    # Generate and print primes up to N
    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)
    
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()