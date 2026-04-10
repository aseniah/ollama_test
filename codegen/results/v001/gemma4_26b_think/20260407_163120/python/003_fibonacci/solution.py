import sys

def main():
    """
    Reads an integer N from the command line and prints the Fibonacci sequence
    up to the largest Fibonacci number that does not exceed N.
    """
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Convert the first command line argument to an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, exit silently
        return

    # If N < 1, the requirement specifies to print nothing
    if n < 1:
        return

    # Initialize the first two numbers of the sequence as specified: 1, 1, ...
    a, b = 1, 1

    # Iterate while the current Fibonacci number does not exceed N
    while a <= n:
        print(a)
        # Calculate the next number in the sequence
        # a becomes the previous b, and b becomes the sum of the previous a and b
        a, b = b, a + b

if __name__ == "__main__":
    main()