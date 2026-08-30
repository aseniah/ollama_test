import sys

def main():
    """
    Reads an integer N from the command line and prints the Fibonacci sequence
    starting from 1, 1, ... up to the largest Fibonacci number that does not exceed N.
    """
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the input argument N
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, the program exits silently 
        # or you could print an error message to stderr.
        return

    # Condition: If N < 1, print nothing.
    if n < 1:
        return

    # Initialize the first two numbers of the Fibonacci sequence
    a, b = 1, 1

    # Continue generating numbers as long as the current number does not exceed N
    while a <= n:
        print(a)
        # Update sequence: next number is the sum of the previous two
        a, b = b, a + b

if __name__ == "__main__":
    main()