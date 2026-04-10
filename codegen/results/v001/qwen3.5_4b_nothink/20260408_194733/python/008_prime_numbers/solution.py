import sys

def is_prime(n):
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(n ** 0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    try:
        # Get integer N from command line arguments
        if len(sys.argv) != 2:
            print("Usage: python script.py <integer_N>", file=sys.stderr)
            sys.exit(1)
        
        n = int(sys.argv[1])
        
        if n < 2:
            # No primes less than 2
            return
        
        # Generate and print all primes up to and including n
        for num in range(2, n + 1):
            if is_prime(num):
                print(num)
                
    except ValueError:
        print("Error: N must be an integer.", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()