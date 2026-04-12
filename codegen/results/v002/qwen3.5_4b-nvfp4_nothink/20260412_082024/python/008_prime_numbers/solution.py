def is_prime(n):
    """Check if a number n is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    i = 3
    while i * i <= n:
        if n % i == 0:
            return False
        i += 2
    return True

def main():
    try:
        # Accept a single integer argument N from the command line
        import sys
        args = sys.argv[1:]
        if not args:
            print("Usage: python script.py <N>")
            sys.exit(1)
        
        n = int(args[0])
    except Exception as e:
        print(f"Error: Invalid input. {e}")
        sys.exit(1)

    # Find and print all prime numbers up to and including N
    primes = []
    for num in range(2, n + 1):
        if is_prime(num):
            primes.append(num)
    
    # Print each prime on a new line
    for prime in primes:
        print(prime)

if __name__ == "__main__":
    main()