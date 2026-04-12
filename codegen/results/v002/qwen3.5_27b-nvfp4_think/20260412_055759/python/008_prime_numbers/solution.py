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
    # Check if exactly one command line argument is provided
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # Invalid integer input, do nothing
        return

    # Print all prime numbers up to and including n
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()