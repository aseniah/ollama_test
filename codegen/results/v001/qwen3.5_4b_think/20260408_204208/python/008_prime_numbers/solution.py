import sys

def is_prime(num):
    """Check if a number is prime."""
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    for i in range(3, int(num ** 0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

def main():
    # Accept a single integer argument N from the command line
    if len(sys.argv) < 2:
        sys.exit(1)
        
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)

    # Print all prime numbers up to and including N, one per line
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()