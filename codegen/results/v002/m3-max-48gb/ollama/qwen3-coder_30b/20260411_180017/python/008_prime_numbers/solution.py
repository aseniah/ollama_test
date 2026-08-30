import sys

def is_prime(n):
    """Check if a number is prime"""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    
    # Check odd divisors up to sqrt(n)
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def main():
    # Read the command line argument
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer")
        sys.exit(1)
    
    # Print all primes up to and including N
    for i in range(2, N + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()