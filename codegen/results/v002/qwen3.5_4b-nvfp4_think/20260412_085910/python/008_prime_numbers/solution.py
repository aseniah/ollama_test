import sys

def is_prime(num):
    """Check if a number is prime using trial division."""
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    
    # Check odd divisors from 3 up to sqrt(num)
    for i in range(3, int(num**0.5) + 1, 2):
        if num % i == 0:
            return False
    
    return True

def main():
    """Process command-line arguments and print primes."""
    if len(sys.argv) != 2:
        print("Usage:", sys.argv[0], "N", file=sys.stderr)
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: argument must be an integer", file=sys.stderr)
        return
    
    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)
    
    # Print all prime numbers (or nothing if none exist)
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()