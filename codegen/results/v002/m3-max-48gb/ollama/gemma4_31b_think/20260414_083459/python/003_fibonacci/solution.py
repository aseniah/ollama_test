import sys

def main():
    """
    Prints the Fibonacci sequence up to and including the largest 
    Fibonacci number that does not exceed the provided integer N.
    """
    # Ensure a command line argument is provided
    if len(sys.argv) < 2:
        return

    try:
        # Convert the argument to an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, we exit silently
        return

    # If N < 1, the problem specifies to print nothing
    if n < 1:
        return

    # Initialize the first two numbers of the sequence
    a, b = 1, 1

    # Print the first number of the sequence
    print(a)

    # Loop and print subsequent Fibonacci numbers while they do not exceed N
    while b <= n:
        print(b)
        # Update Fibonacci sequence: next number is sum of previous two
        a, b = b, a + b

if __name__ == "__main__":
    main()