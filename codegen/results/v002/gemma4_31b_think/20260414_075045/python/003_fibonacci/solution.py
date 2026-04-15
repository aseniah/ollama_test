import sys

def main():
    """
    Prints Fibonacci numbers starting from 1, 1, 2, ... 
    up to and including the largest Fibonacci number that does not exceed N.
    """
    # Ensure that an argument is provided via the command line
    if len(sys.argv) < 2:
        return

    try:
        # Parse the command line argument N as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, we terminate silently
        return

    # Per requirements: If N < 1, print nothing.
    if n < 1:
        return

    # Starting numbers of the sequence as defined: 1, 1, 2, 3, 5...
    a, b = 1, 1

    # Print the first number of the sequence
    print(a)

    # Generate and print the sequence as long as the number does not exceed N
    while b <= n:
        print(b)
        # Calculate next Fibonacci number: F(n) = F(n-1) + F(n-2)
        a, b = b, a + b

if __name__ == "__main__":
    main()