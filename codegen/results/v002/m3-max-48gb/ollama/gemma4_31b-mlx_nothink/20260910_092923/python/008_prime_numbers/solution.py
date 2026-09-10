import sys

def is_prime(num):
    """Checks if a number is prime."""
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    # Check odd numbers up to the square root of num
    for i in range(3, int(num**0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

def main():
    # Ensure an argument is provided
    if len(sys.argv) != 2:
        return

    try:
        # Convert the command line argument to an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not an integer, the program exits silently
        return

    # Iterate through all numbers from 2 up to and including N
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()