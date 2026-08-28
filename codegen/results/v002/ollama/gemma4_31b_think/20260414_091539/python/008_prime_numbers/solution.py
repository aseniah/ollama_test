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
    # Ensure a command line argument was provided
    if len(sys.argv) != 2:
        # The prompt specifies accepting a single integer argument.
        # If none is provided or too many are provided, we exit silently 
        # or could print a usage message.
        return

    try:
        # Convert the argument to an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, print nothing/exit
        return

    # Iterate through all numbers up to and including N
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()