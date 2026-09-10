import sys

def is_prime(n):
    """Checks if a number is prime."""
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    # Check odd divisors up to the square root of n
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

def main():
    # Ensure an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the first command line argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, exit silently or handle error
        return

    # Iterate from 2 up to and including N
    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()