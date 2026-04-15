import sys

def main():
    # Check if an argument was provided
    if len(sys.argv) < 2:
        return

    try:
        # Parse the input N from the command line argument
        n = int(sys.argv[1])
    except ValueError:
        # If the argument is not an integer, do nothing
        return

    # If N < 1, print nothing as per requirements
    if n < 1:
        return

    # Initialize the first two Fibonacci numbers
    a, b = 1, 1

    # Print the first number if it doesn't exceed N
    if a <= n:
        print(a)
    else:
        return

    # Loop to generate and print Fibonacci numbers
    # We use a while loop to check the current number against N
    while b <= n:
        print(b)
        # Update values: a becomes the previous b, b becomes the sum
        a, b = b, a + b

if __name__ == "__main__":
    main()