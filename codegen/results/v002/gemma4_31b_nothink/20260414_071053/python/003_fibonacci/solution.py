import sys

def main():
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the input N from the command line
        n = int(sys.argv[1])
    except ValueError:
        # If the input is not a valid integer, we exit silently or could print an error
        return

    # If N < 1, the requirements state to print nothing
    if n < 1:
        return

    # Initialize the first two Fibonacci numbers
    a, b = 1, 1

    # Print the first number in the sequence
    print(a)

    # Continue generating and printing numbers as long as they do not exceed N
    # We start by checking if the second '1' is within bounds
    while b <= n:
        print(b)
        # Calculate the next Fibonacci number
        a, b = b, a + b

if __name__ == "__main__":
    main()