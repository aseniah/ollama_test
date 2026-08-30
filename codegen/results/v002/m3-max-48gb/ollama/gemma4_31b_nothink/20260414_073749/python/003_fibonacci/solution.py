import sys

def main():
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the command line argument as an integer
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not a valid integer, do nothing or handle error
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two Fibonacci numbers
    a, b = 1, 1

    # We start by printing the first number of the sequence
    # The sequence is defined as 1, 1, 2, 3, 5...
    while a <= n:
        print(a)
        # Calculate the next number in the sequence
        # a becomes the current, b becomes the next
        a, b = b, a + b

if __name__ == "__main__":
    main()