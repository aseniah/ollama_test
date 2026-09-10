import sys

def is_prime(num):
    """Checks if a number is prime."""
    if num < 2:
        return False
    # Check divisors from 2 up to the square root of the number
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

def main():
    # Ensure there is at least one argument passed via command line
    if len(sys.argv) < 2:
        return

    try:
        # Parse the first command line argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, exit silently
        return

    # Iterate through all numbers from 2 up to and including N
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()