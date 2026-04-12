import sys

def is_prime(n: int) -> bool:
    """Check if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    # Check odd divisors up to sqrt(n)
    limit = int(n ** 0.5) + 1
    for i in range(3, limit, 2):
        if n % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        # If no argument provided, exit silently or handle as needed.
        # The problem implies exactly one argument, so we exit if invalid.
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, exit silently.
        return

    # Print all primes up to and including n
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()